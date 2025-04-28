package seckill

import (
	"context"
	"kitex-server/inits"
	"kitex-server/seckill/dao"
	"kitex-server/seckill/kafka"
	seckill "kitex-server/seckill/kitex_gen/seckill"
	"kitex-server/seckill/model"
	"kitex-server/seckill/seckill_resp"
	"kitex-server/users/kitex_gen/user"
	"kitex-server/utils"
	"time"
)

// SecKillServiceImpl implements the last service interface defined in the IDL.
type SecKillServiceImpl struct{}

// SecKill implements the SecKillServiceImpl interface.
func (s *SecKillServiceImpl) SecKill(ctx context.Context, req *seckill.SecKillRequest) (resp *seckill.SecKillResponse, err error) {
	var emptyStatus seckill.Status
	//1.检查抢购是否开始
	//1.1.获取redis中的秒杀活动信息
	event, status := dao.GetSecKillEventInRedis(req.ItemId)
	if status != emptyStatus {
		return &seckill.SecKillResponse{Status: &status}, nil
	}
	//1.2.检查当前时间是否在秒杀活动的开始时间和结束时间之间
	loc, _ := time.LoadLocation("Asia/Shanghai")
	nowTimeUnix := time.Now().Unix()
	eventStartTimeUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", event.StartTime, loc)
	eventEndTimeUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", event.EndTime, loc)
	//1.3.如果不在范围内，返回秒杀活动未开始或已结束
	if nowTimeUnix < eventStartTimeUnix.Unix() || nowTimeUnix > eventEndTimeUnix.Unix() {
		return &seckill.SecKillResponse{Status: &seckill_resp.NotInSecKillTime}, nil
	}
	//2.进行redis缓存预扣
	status = dao.DeductStockInRedis(req.ItemId)
	if status != emptyStatus {
		return &seckill.SecKillResponse{Status: &status}, nil
	}
	//3.如果预扣成功，将抢购请求放入kafka
	//3.1.生成订单号
	orderID, status := utils.NewOrderIDGenerator(inits.Re, "orderID").GenerateID(ctx)
	if status != emptyStatus {
		return &seckill.SecKillResponse{Status: &status}, nil
	}
	//3.2.将抢购请求放入kafka
	err = kafka.AddMsgToKafka(req.ItemId, orderID)
	if err != nil {
		retErr := seckill_resp.InternalErr(err)
		return &seckill.SecKillResponse{Status: &retErr}, nil
	}
	return &seckill.SecKillResponse{Status: &seckill_resp.Ok, OrderId: &orderID}, nil
}

// CreateSecKill implements the SecKillServiceImpl interface.
func (s *SecKillServiceImpl) CreateSecKill(ctx context.Context, req *seckill.CreateSecKillRequest) (resp *seckill.CreateSecKillResponse, err error) {
	var emptyStatus seckill.Status
	var event model.CreateSecKillEvent
	var emptyUserStatus user.Status
	//1.验证token
	_, userStatus := utils.CheckJwtToken(req.Token)
	if userStatus != emptyUserStatus {
		return &seckill.CreateSecKillResponse{Status: &seckill_resp.InvalidToken}, nil
	}
	//2.创建秒杀活动
	event.ItemID = req.ItemId
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t1 := time.Unix(*req.StartTime, 0).In(loc)
	t2 := time.Unix(*req.EndTime, 0).In(loc)
	timeStr1 := t1.Format("2006-01-02 15:04:05")
	timeStr2 := t2.Format("2006-01-02 15:04:05")
	event.StartTime = timeStr1
	event.EndTime = timeStr2
	//将请求中的秒杀活动信息存入redis
	status := dao.PreHeatStockToRedis(event)
	if status != emptyStatus {
		return &seckill.CreateSecKillResponse{Status: &status}, nil
	}
	return &seckill.CreateSecKillResponse{Status: &seckill_resp.Ok}, nil
}

// GetOrderStatus implements the SecKillServiceImpl interface.
func (s *SecKillServiceImpl) GetOrderStatus(ctx context.Context, req *seckill.GetOrderStatusRequest) (resp *seckill.GetOrderStatusResponse, err error) {
	var emptyStatus seckill.Status
	orderStatus, status := dao.GetOrderStatusInMysql(req.OrderId)
	if status != emptyStatus {
		return &seckill.GetOrderStatusResponse{Status: &status}, nil
	}

	return &seckill.GetOrderStatusResponse{
		Status:      &seckill_resp.Ok,
		OrderStatus: &orderStatus.Status,
		OrderId:     &orderStatus.OrderNumber,
		ItemName:    &orderStatus.ProductID,
	}, nil
}
