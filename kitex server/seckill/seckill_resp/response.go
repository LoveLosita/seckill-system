package seckill_resp

import "kitex-server/seckill/kitex_gen/seckill"

func InternalErr(err error) seckill.Status {
	return seckill.Status{Code: "500", Message: err.Error()}
}

var (
	Ok = seckill.Status{ //正常,和客户端统一
		Code:    "10000",
		Message: "ok",
	}
	ItemNotFound = seckill.Status{ //商品不存在
		Code:    "42001",
		Message: "item not found",
	}
	OutOfStock = seckill.Status{ //商品库存不足
		Code:    "42002",
		Message: "out of stock",
	}
)
