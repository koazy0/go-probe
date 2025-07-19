package main

import "github.com/koazy0/go-probe/monitor"
import _ "github.com/koazy0/go-probe/global"

func main() {

	monitor.FM.StartMonitoring([]string{"test", "test"})
	monitor.FM.StopMonitoring()
	monitor.FM.Events()
}
