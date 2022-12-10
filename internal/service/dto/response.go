package dto

import "time"

type GetLimitResponse struct {
	MaxItemCount  uint64        `json:"max_item_count"`
	ProcessPeriod time.Duration `json:"process_period"`
}
