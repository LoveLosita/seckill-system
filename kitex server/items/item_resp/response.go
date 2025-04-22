package item_resp

import (
	"kitex-server/items/kitex_gen/items"
)

func InternalErr(err error) items.Status {
	return items.Status{Code: "500", Message: err.Error()}
}

var (
	Ok = items.Status{ //正常,和客户端统一
		Code:    "10000",
		Message: "ok",
	}
	ItemNotFound = items.Status{ //物品不存在
		Code:    "41001",
		Message: "item not found",
	}
)
