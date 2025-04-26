package model

type CreateSecKillEvent struct {
	ItemID    int64  `json:"item_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Stock     int64  `json:"stock"`
}

type Order struct {
	OrderNumber string `json:"order_number"`
	ProductName string `json:"product_name"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
