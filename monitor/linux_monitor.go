//go:build linux

package monitor

import (
	"github.com/koazy0/go-probe/model"
)

type LinuxMonitor struct {
}

func newFileMonitor() FileMonitor {
	return &LinuxMonitor{}
}

func (l *LinuxMonitoring) StartMonitoring(paths []string) error {
	zap.S().Infof("Linux StartMonitoring called")
	return nil
}

func (l *LinuxMonitoring) StopMonitoring() error {
	zap.S().Infof("Linux StopMonitoring called")
	return nil
}

func (l *LinuxMonitoring) Events() <-chan model.FileEvent {
	zap.S().Infof("Linux Events called")
	return nil
}
