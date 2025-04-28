package model

type JsonSecKillEvent struct {
	ItemId    int64  `json:"item_id"`
	Amount    int64  `json:"amount,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	Token     string `json:"token"`
}
