package model

import "time"

type Item struct {
	Id        int
	Name      string
	Price     float64
	Stock     int
	Intro     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
