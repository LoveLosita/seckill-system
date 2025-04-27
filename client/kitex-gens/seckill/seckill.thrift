namespace go seckill

struct Status {
  1: required string code
  2: required string message
}

struct SecKillRequest { #秒杀请求
  1: required i64 item_id
}

struct SecKillResponse { #秒杀响应
  1: required Status status
  2: optional string order_id
}

struct CreateSecKillRequest { #创建秒杀请求
  1: required i64 item_id
  2: optional i64 amount
  3: optional i64 start_time
  4: optional i64 end_time
  5: required string token
}

struct CreateSecKillResponse { #创建秒杀响应
  1: required Status status
}

struct GetOrderStatusRequest { #获取订单状态请求
  1: required string order_id
}

struct GetOrderStatusResponse { #获取订单状态响应
  1: required Status status
  2: optional string order_id
  3: optional string item_name
  4: optional string order_status
}

service SecKillService {
  SecKillResponse sec_kill(1: SecKillRequest req) #秒杀
  CreateSecKillResponse create_sec_kill(1: CreateSecKillRequest req) #创建秒杀
  GetOrderStatusResponse get_order_status(1: GetOrderStatusRequest req) #获取订单状态
}