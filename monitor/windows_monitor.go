package monitor

import (
	"github.com/koazy0/go-probe/model"
	"go.uber.org/zap"
)

type WindowsMonitor struct {
}

func newFileMonitor() FileMonitor {
	return &WindowsMonitor{}
}
func (w *WindowsMonitor) StartMonitoring(paths []string) error {
	zap.S().Infof("Windows StartMonitoring called")
	return nil
}

func (w *WindowsMonitor) StopMonitoring() error {
	zap.S().Infof("Windows StopMonitoring called")
	return nil
}

func (w *WindowsMonitor) Events() <-chan model.FileEvent {
	zap.S().Infof("Windows Events called")
	return nil
}
