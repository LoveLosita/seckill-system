package dao

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"kitex-server/inits"
	model2 "kitex-server/items/model"
	"kitex-server/seckill/kitex_gen/seckill"
	"kitex-server/seckill/model"
	"kitex-server/seckill/seckill_resp"
	"strconv"
)

func PreHeatStockToRedis(req model.CreateSecKillEvent) seckill.Status {
	var emptyStatus seckill.Status
	ctx := context.Background()
	if req.Stock == 0 { //如果没有传入库存，则从数据库中获取
		stock, status := GetItemStock(req.ItemID)
		if status != emptyStatus {
			return status
		}
		req.Stock = stock
	}
	//将秒杀活动的开始时间、结束时间、库存等信息存入redis
	hashKey := fmt.Sprintf("seckill:event:%d", req.ItemID)
	hashData := map[string]interface{}{
		"item_name":  req.ItemID,
		"start_time": req.StartTime,
		"end_time":   req.EndTime,
		"stock":      req.Stock, //会以字符串存储
	}
	if err := inits.Re.HSet(ctx, hashKey, hashData).Err(); err != nil {
		return seckill_resp.InternalErr(err)
	}
	return seckill.Status{}
}

func GetItemStock(itemID int64) (int64, seckill.Status) {
	var item model2.Item
	result := inits.Db.Table("items").Where("id = ?", itemID).Select("stock").First(&item)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return -1, seckill_resp.ItemNotFound
		}
		return -1, seckill_resp.InternalErr(result.Error)
	}
	return item.Stock, seckill.Status{}
}

func DeductStockInMysql(itemID int64) seckill.Status {
	//1.先看看该商品是否已经在表单中
	var item model2.Item
	result := inits.Db.Table("items").Where("id = ?", itemID).First(&item)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return seckill_resp.ItemNotFound
		} else {
			return seckill_resp.InternalErr(result.Error)
		}
	}
	//2.如果在，则将库存减一
	item.Stock--
	result = inits.Db.Table("items").Where("id = ?", itemID).Updates(&item)
	if result.Error != nil {
		return seckill_resp.InternalErr(result.Error)
	}
	return seckill.Status{}
}

func DeductStockInRedis(itemID int64) seckill.Status {
	//1.先看看该商品是否已经在缓存中
	ctx := context.Background()
	stock, err := inits.Re.HGet(ctx, fmt.Sprintf("seckill:event:%d", itemID), "stock").Result()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return seckill_resp.ItemNotFound
		} else {
			return seckill_resp.InternalErr(err)
		}
	}
	if stock == "" { //如果没有，则返回错误
		return seckill_resp.ItemNotFound
	}
	//2.如果有，则将库存减一
	// 模拟原子扣减库存的操作，使用 HINCRBY 命令进行原子递减
	hashKey := fmt.Sprintf("seckill:event:%d", itemID)
	field := "stock"
	decrement := int64(1) // 每次扣减 1
	newStock, err := inits.Re.HIncrBy(ctx, hashKey, field, -decrement).Result()
	if err != nil {
		seckill_resp.InternalErr(fmt.Errorf("扣减库存失败: %v", err))
	}
	if newStock < 0 {
		return seckill_resp.OutOfStock
	}
	return seckill.Status{}
}

func AddOrderStatusToMysql(orderID string, productID int64, status string) seckill.Status {
	var order model.Order
	order.OrderNumber = orderID
	order.ProductName = strconv.FormatInt(productID, 10)
	order.Status = status
	result := inits.Db.Table("orders").Create(&order)
	if result.Error != nil {
		return seckill_resp.InternalErr(result.Error)
	}
	return seckill.Status{}
}

func GetSecKillEventInRedis(itemID int64) (model.CreateSecKillEvent, seckill.Status) {
	var event model.CreateSecKillEvent
	ctx := context.Background()
	hashKey := fmt.Sprintf("seckill:event:%d", itemID)
	//从redis中获取秒杀活动信息
	hashData, err := inits.Re.HGetAll(ctx, hashKey).Result()
	if err != nil {
		return event, seckill_resp.InternalErr(err)
	}
	if len(hashData) == 0 {
		return event, seckill_resp.ItemNotFound
	}
	event.ItemID = itemID
	event.StartTime = hashData["start_time"]
	event.EndTime = hashData["end_time"]
	event.Stock, _ = strconv.ParseInt(hashData["stock"], 10, 64)
	return event, seckill.Status{}
}

func GetOrderStatusInMysql(orderID string) (model.Order, seckill.Status) {
	var order model.Order
	result := inits.Db.Table("orders").Where("order_number = ?", orderID).First(&order)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return order, seckill_resp.OrderNotFound
		} else {
			return order, seckill_resp.InternalErr(result.Error)
		}
	}
	return order, seckill.Status{}
}
