package dao

import (
	"kitex-server/inits"
	"kitex-server/items/item_resp"
	"kitex-server/items/kitex_gen/items"
	"kitex-server/items/model"
)

func InsertItem(item model.Item) items.Status {
	result := inits.Db.Create(&item)
	if result.Error != nil {
		return item_resp.InternalErr(result.Error)
	}
	return items.Status{}
}
