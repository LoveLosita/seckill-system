package items

import (
	"context"
	"kitex-server/items/dao"
	"kitex-server/items/item_resp"
	items "kitex-server/items/kitex_gen/items"
	"kitex-server/items/model"
	"kitex-server/users/kitex_gen/user"
	"kitex-server/utils"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetItemInfo implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItemInfo(ctx context.Context, req *items.GetItemInfoRequest) (resp *items.GetItemInfoResponse, err error) {
	var emptyStatus items.Status
	//1.获取商品信息
	item, status := dao.GetItemByID(req.Id)
	if status != emptyStatus {
		return &items.GetItemInfoResponse{Status: &status}, nil
	}
	var retItem items.Item
	//类型转换
	retItem.Id = &item.Id
	retItem.Name = &item.Name
	retItem.Price = &item.Price
	retItem.Stock = &item.Stock
	retItem.Intro = &item.Intro
	unixCrAt := item.CreatedAt.Unix()
	unixUpAt := item.UpdatedAt.Unix()
	retItem.CreatedAt = &unixCrAt
	retItem.UpdatedAt = &unixUpAt
	//2.返回结果
	return &items.GetItemInfoResponse{
		Status: &item_resp.Ok,
		Item:   &retItem,
	}, nil
}

// GetItemList implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItemList(ctx context.Context, req *items.GetItemListRequest) (resp *items.GetItemListResponse, err error) {
	var emptyStatus items.Status
	itemList, status := dao.GetAllItems()
	if status != emptyStatus {
		return &items.GetItemListResponse{Status: &status}, nil
	}
	var ptrItemList []*items.Item
	for i := range itemList {
		//类型转换
		var retItem items.Item
		retItem.Id = &itemList[i].Id
		retItem.Name = &itemList[i].Name
		retItem.Price = &itemList[i].Price
		retItem.Stock = &itemList[i].Stock
		retItem.Intro = &itemList[i].Intro
		unixCrAt := itemList[i].CreatedAt.Unix()
		unixUpAt := itemList[i].UpdatedAt.Unix()
		retItem.CreatedAt = &unixCrAt
		retItem.UpdatedAt = &unixUpAt
		ptrItemList = append(ptrItemList, &retItem)
	}
	return &items.GetItemListResponse{
		Status: &item_resp.Ok,
		Items:  ptrItemList,
	}, nil
}

// AddItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) AddItem(ctx context.Context, req *items.AddItemRequest) (resp *items.AddItemResponse, err error) {
	var emptyStatus items.Status
	var emptyUserStatus user.Status
	//1.先获取token
	if req != nil && req.Token != "" {
		itemID, userStatus := utils.CheckJwtToken(req.Token)
		if itemID == 0 || userStatus != emptyUserStatus {
			return &items.AddItemResponse{Status: (*items.Status)(&userStatus)}, nil
		}
	} else if req == nil {
		return &items.AddItemResponse{Status: &item_resp.EmptyRequest}, nil
	} else {
		return &items.AddItemResponse{Status: &item_resp.NotLoggedIn}, nil
	}
	//2.再插入商品
	var insertItem model.Item
	insertItem.Name = req.Name
	insertItem.Price = req.Price
	insertItem.Stock = req.Stock
	insertItem.Intro = req.Intro
	status := dao.InsertItem(insertItem)
	if status != emptyStatus {
		return &items.AddItemResponse{Status: &status}, nil
	}
	return &items.AddItemResponse{Status: &item_resp.Ok}, nil
}

// UpdateItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) UpdateItem(ctx context.Context, req *items.UpdateItemRequest) (resp *items.UpdateItemResponse, err error) {
	var emptyStatus items.Status
	var emptyUserStatus user.Status
	//1.先获取token
	if req != nil && req.Token != "" {
		itemID, userStatus := utils.CheckJwtToken(req.Token)
		if itemID == 0 || userStatus != emptyUserStatus {
			return &items.UpdateItemResponse{Status: (*items.Status)(&userStatus)}, nil
		}
	} else if req == nil {
		return &items.UpdateItemResponse{Status: &item_resp.EmptyRequest}, nil
	} else {
		return &items.UpdateItemResponse{Status: &item_resp.NotLoggedIn}, nil
	}
	//2.更新商品
	var updateItem model.Item
	updateItem.Id = *req.Id
	updateItem.Name = *req.Name
	updateItem.Price = *req.Price
	updateItem.Stock = *req.Stock
	updateItem.Intro = *req.Intro
	status := dao.UpdateItem(updateItem)
	if status != emptyStatus {
		return &items.UpdateItemResponse{Status: &status}, nil
	}
	return &items.UpdateItemResponse{Status: &item_resp.Ok}, nil
}

// DeleteItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) DeleteItem(ctx context.Context, req *items.DeleteItemRequest) (resp *items.DeleteItemResponse, err error) {
	var emptyStatus items.Status
	var emptyUserStatus user.Status
	//1.先获取token
	if req != nil && req.Token != "" {
		itemID, userStatus := utils.CheckJwtToken(req.Token)
		if itemID == 0 || userStatus != emptyUserStatus {
			return &items.DeleteItemResponse{Status: (*items.Status)(&userStatus)}, nil
		}
	} else if req == nil {
		return &items.DeleteItemResponse{Status: &item_resp.EmptyRequest}, nil
	} else {
		return &items.DeleteItemResponse{Status: &item_resp.NotLoggedIn}, nil
	}
	//2.删除商品
	status := dao.DeleteItem(req.Id)
	if status != emptyStatus {
		return &items.DeleteItemResponse{Status: &status}, nil
	}
	return &items.DeleteItemResponse{Status: &item_resp.Ok}, nil
}
