package model

type Job struct {
	ID     int64 `json:"id"`
	Input  int64 `json:"input"`
	Result int64 `json:"result"`
}
