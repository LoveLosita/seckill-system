package dao

import (
	"errors"
	"gorm.io/gorm"
	"kitex-server/inits"
	"kitex-server/items/item_resp"
	"kitex-server/items/kitex_gen/items"
	"kitex-server/items/model"
)

func InsertItem(item model.Item) items.Status {
	result := inits.Db.Table("items").Create(&item)
	if result.Error != nil {
		return item_resp.InternalErr(result.Error)
	}
	return items.Status{}
}

func GetItemByID(itemID int64) (model.Item, items.Status) {
	var item model.Item
	result := inits.Db.Table("items").First(&item, "id = ?", itemID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Item{}, item_resp.ItemNotFound
		}
		return model.Item{}, item_resp.InternalErr(result.Error)
	}
	return item, items.Status{}
}

func UpdateItem(item model.Item) items.Status {
	result := inits.Db.Table("items").Where("id = ?", item.Id).Updates(&item)
	if result.Error != nil {
		return item_resp.InternalErr(result.Error)
	}
	return items.Status{}
}

func DeleteItem(itemID int64) items.Status {
	result := inits.Db.Table("items").Delete(&model.Item{}, itemID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return item_resp.ItemNotFound
		}
		return item_resp.InternalErr(result.Error)
	}
	return items.Status{}
}

func IfItemExists(itemID int64) (bool, items.Status) {
	var item model.Item
	result := inits.Db.Table("items").First(&item, "id = ?", itemID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, items.Status{}
		}
		return false, item_resp.InternalErr(result.Error)
	}
	return true, items.Status{}
}

func GetAllItems() ([]model.Item, items.Status) {
	var itemsList []model.Item
	result := inits.Db.Table("items").Find(&itemsList)
	if result.Error != nil {
		return nil, item_resp.InternalErr(result.Error)
	}
	return itemsList, items.Status{}
}
