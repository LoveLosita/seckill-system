package items

import (
	"context"
	items "kitex-server/items/kitex_gen/items"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetItemInfo implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItemInfo(ctx context.Context, req *items.GetItemInfoRequest) (resp *items.GetItemInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetItemList implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItemList(ctx context.Context, req *items.GetItemListRequest) (resp *items.GetItemListResponse, err error) {
	// TODO: Your code here...
	return
}

// AddItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) AddItem(ctx context.Context, req *items.AddItemRequest) (resp *items.AddItemResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) UpdateItem(ctx context.Context, req *items.UpdateItemRequest) (resp *items.UpdateItemResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) DeleteItem(ctx context.Context, req *items.DeleteItemRequest) (resp *items.DeleteItemResponse, err error) {
	// TODO: Your code here...
	return
}
