package model

type Job struct {
	ID     uint64 `json:"id"`
	Input  uint64 `json:"input"`
	Result int64  `json:"result"`
}
