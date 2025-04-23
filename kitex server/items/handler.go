package items

import (
	"context"
	"kitex-server/items/dao"
	"kitex-server/items/item_resp"
	items "kitex-server/items/kitex_gen/items"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetItemInfo implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItemInfo(ctx context.Context, req *items.GetItemInfoRequest) (resp *items.GetItemInfoResponse, err error) {
	var emptyStatus items.Status
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
	//返回结果
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
		ptrItemList = append(ptrItemList, &itemList[i])
	}
	return &items.GetItemListResponse{
		Status: &item_resp.Ok,
		Items:  ptrItemList,
	}, nil
}

// AddItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) AddItem(ctx context.Context, req *items.AddItemRequest) (resp *items.AddItemResponse, err error) {
	var emptyStatus items.Status
	status := dao.InsertItem(*req)
	if status != emptyStatus {
		return &items.AddItemResponse{Status: &status}, nil
	}
	return &items.AddItemResponse{Status: &item_resp.Ok}, nil
}

// UpdateItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) UpdateItem(ctx context.Context, req *items.UpdateItemRequest) (resp *items.UpdateItemResponse, err error) {
	var emptyStatus items.Status
	status := dao.UpdateItem(*req)
	if status != emptyStatus {
		return &items.UpdateItemResponse{Status: &status}, nil
	}
	return &items.UpdateItemResponse{Status: &item_resp.Ok}, nil
}

// DeleteItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) DeleteItem(ctx context.Context, req *items.DeleteItemRequest) (resp *items.DeleteItemResponse, err error) {
	var emptyStatus items.Status
	status := dao.DeleteItem(req.Id)
	if status != emptyStatus {
		return &items.DeleteItemResponse{Status: &status}, nil
	}
	return &items.DeleteItemResponse{Status: &item_resp.Ok}, nil
}
