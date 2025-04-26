package seckill

import (
	"context"
	"kitex-server/seckill/dao"
	seckill "kitex-server/seckill/kitex_gen/seckill"
	"kitex-server/seckill/model"
	"kitex-server/seckill/seckill_resp"
	"time"
)

// SecKillServiceImpl implements the last service interface defined in the IDL.
type SecKillServiceImpl struct{}

// SecKill implements the SecKillServiceImpl interface.
func (s *SecKillServiceImpl) SecKill(ctx context.Context, req *seckill.SecKillRequest) (resp *seckill.SecKillResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateSecKill implements the SecKillServiceImpl interface.
func (s *SecKillServiceImpl) CreateSecKill(ctx context.Context, req *seckill.CreateSecKillRequest) (resp *seckill.CreateSecKillResponse, err error) {
	var emptyStatus seckill.Status
	var event model.CreateSecKillEvent
	event.ItemID = req.ItemId
	t1 := time.Unix(*req.StartTime, 0)
	t2 := time.Unix(*req.EndTime, 0)
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
	// TODO: Your code here...
	return
}
