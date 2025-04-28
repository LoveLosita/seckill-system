package model

import "time"

type CreateSecKillEvent struct {
	ItemID    int64  `json:"item_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Stock     int64  `json:"stock"`
}

type Order struct {
	OrderNumber string    `json:"order_number"`
	ProductID   string    `json:"product_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}
