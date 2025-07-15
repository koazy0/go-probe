package main

import (
	"context"
	"io"
	"log"
	"net"
	"sync"
)

type proxyRule struct {
	SourceAddr      string
	DestinationAddr string
	ctx             context.Context
}

var rules = []proxyRule{
	{":10030", "127.0.0.1:22", nil},
	{":10040", "127.0.0.1:22", nil},
}

func main() {
	dynamicRuleListen()

	for _, rule := range rules {
		listener, err := net.Listen("tcp", rule.SourceAddr)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go func(listener net.Listener, dstAddr string) {
			for {
				srcConn, err := listener.Accept()
				if err != nil {
					return
				}
				go handleConnection(srcConn, dstAddr)
			}
		}(listener, rule.DestinationAddr)

	}

	select {}
}

func handleConnection(srcConn net.Conn, dstAddr string) {
	defer srcConn.Close()
	dstConn, err := net.Dial("tcp", dstAddr)
	if err != nil {
		return
	}
	defer dstConn.Close()
	wg := sync.WaitGroup{}
	wg.Add(2)

	//转发访问本机的流量
	go func() {
		io.Copy(dstConn, srcConn)
		wg.Done()
	}()
	//转发返回的流量
	go func() {
		io.Copy(srcConn, dstConn)
		wg.Done()
	}()
	wg.Wait()
}

// 获取动态下发功能
func dynamicRuleListen() {
	listener, err := net.Listen("tcp", "50050")
	if err != nil {
		return
	}
	listener.Accept()
}
