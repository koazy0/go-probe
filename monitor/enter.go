//go:build windows

package monitor

import (
	"github.com/koazy0/go-probe/model"
)

type FileMonitor interface {
	StartMonitoring(paths []string) error
	StopMonitoring() error
	Events() <-chan model.FileEvent
}

var FM FileMonitor

func init() {
	FM = newFileMonitor()
	//判断当前是哪个平台从而返回对应的FileMonitor
}
