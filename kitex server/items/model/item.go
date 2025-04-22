package model

import "time"

type Item struct {
	Id        int64
	Name      string
	Price     float64
	Stock     int64
	Intro     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
