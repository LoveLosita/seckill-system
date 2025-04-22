namespace go items

struct Status {
  1: required string code
  2: required string message
}

struct Item {
  1: optional i64 id
  2: optional string name
  3: optional double price
  4: optional i64 stock
  5: optional string intro
  6: optional i64 created_at
  7: optional i64 updated_at
}

struct GetItemInfoRequest {
  1: required i64 id
}

struct GetItemInfoResponse {
  1: required Status status
  2: optional Item item
}

struct GetItemListRequest {
}

struct GetItemListResponse {
  1: required Status status
  2: optional list<Item> items
}

struct AddItemRequest {
  1: required string name
  2: required double price
  3: required i64 stock
  4: required string intro
}

struct AddItemResponse {
  1: required Status status
  2: optional Item item
}

struct UpdateItemRequest {
  1: optional i64 id
  2: optional string name
  3: optional double price
  4: optional i64 stock
  5: optional string intro
}

struct UpdateItemResponse {
  1: required Status status
  2: optional Item item
}

struct DeleteItemRequest {
  1: required i64 id
}

struct DeleteItemResponse {
  1: required Status status
}

service ItemService {
  GetItemInfoResponse get_item_info(1: GetItemInfoRequest req) #获取商品信息
  GetItemListResponse get_item_list(1: GetItemListRequest req) #获取商品列表
  AddItemResponse add_item(1: AddItemRequest req) #添加商品
  UpdateItemResponse update_item(1: UpdateItemRequest req) #更新商品
  DeleteItemResponse delete_item(1: DeleteItemRequest req) #删除商品
}