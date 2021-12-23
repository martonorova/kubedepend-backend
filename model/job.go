package model

type Job struct {
	ID     uint64 `json:"id" gorm:"primaryKey"`
	Input  uint64 `json:"input"`
	Result int64  `json:"result"`
}
