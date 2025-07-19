package model

import "time"

type FileEvent struct {
	FilePath  string
	EventType string // create / write / read / delete
	Process   string // 访问者进程名
	PID       int
	Timestamp time.Time
}
