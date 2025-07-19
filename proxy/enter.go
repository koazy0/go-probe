package proxy

import (
	"context"
	"github.com/koazy0/go-probe/model"
	"go.uber.org/zap"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
)

var (
	ProxyChan = make(chan []model.ProxyRule)
	//sessionMap = make(map[string]session)
	sessionMap = sync.Map{} // 并发安全的会话映射：key=SrcAddr, value=*session
)

type session struct {
	SrcAddr  string
	DstAddr  string
	Ctx      context.Context
	Cancel   context.CancelFunc
	Listener net.Listener
	Counter  atomic.Int64 //用于计数,couner.load
}

func SetUp() {
	//这里起一个协程进行接收规则
	go func() {
		for newRules := range ProxyChan {
			for _, rule := range newRules {
				//先遍历Map，看看存不存在
				sessionExistAny, ok := sessionMap.Load(rule.SourceAddr)
				switch rule.Option {
				case "add":
					//看看map里面存不存在
					if ok {
						sessionExist, _ := sessionExistAny.(*session)
						//规则不变，不需要作出改动
						if sessionExist.DstAddr == rule.DestinationAddr {
							continue
						}
						//map里面存在，规则改变，直接把原来这个断开
						sessionExist.Stop()
						sessionMap.Delete(rule.SourceAddr)
					}
					//如果不存在或者改动，直接创建一个新的会话Session就
					ctx, cancle := context.WithCancel(context.Background())
					s := &session{
						SrcAddr: rule.SourceAddr,
						DstAddr: rule.DestinationAddr,
						Ctx:     ctx,
						Cancel:  cancle,
					}
					sessionMap.Store(rule.SourceAddr, s)
					go s.Start()
				case "del":
					if ok {
						sessionExist, _ := sessionExistAny.(*session)
						sessionExist.Stop()
						sessionMap.Delete(rule.SourceAddr)
					} else {
						zap.S().Warnf("deleted rule dosen't exist , rule: %#v", rule)
					}
				default:
					zap.S().Errorf("option invalid,rule: %#v", rule)
				}

			}
		}
	}()

}

// Start  负责管理会话连接
func (s *session) Start() {
	//先打开本机端口
	listener, err := net.Listen("tcp", s.SrcAddr)
	if err != nil {
		log.Println(err.Error())
	}
	s.Listener = listener
	//defer listener.Close() //不在这进行close，要立刻关掉，否则会端口冲突
	//为每一个新连接到的端口进行转发
	for {
		srcConn, err := listener.Accept()
		if err != nil {
			zap.S().Error("listener.Accept() error:" + err.Error())
			return
		}
		go s.handleConnection(srcConn, s.DstAddr)
	}
}

// Stop 为会话提供一个cancle的入口
func (s *session) Stop() {
	s.Cancel()
	s.Listener.Close()
}

// handleConnection 进行流量转发
func (s *session) handleConnection(srcConn net.Conn, dstAddr string) {
	s.Counter.Add(1)
	defer func() {
		s.Counter.Add(-1)
		srcConn.Close()
	}()
	connCtx, cancel := context.WithCancel(s.Ctx)
	zap.S().Infof("src addr: %s, use port %s, dst addr: %s, now count of connection %d", srcConn.RemoteAddr().String(), srcConn.LocalAddr().String(), dstAddr, s.Counter.Load())

	//启动连接
	dstConn, err := net.Dial("tcp", dstAddr)
	if err != nil {
		cancel() //及时取消掉，防止上下文泄露
		return
	}
	defer func() {
		dstConn.Close()
	}()

	//转发访问本机的流量
	go func() {
		_, err := io.Copy(dstConn, srcConn)
		if err != nil {
			zap.S().Warnf("connection error: %s", err)
		}
		cancel()
	}()

	//转发返回的流量
	go func() {
		_, err := io.Copy(srcConn, dstConn)
		if err != nil {
			zap.S().Warnf("connection error: %s", err)
		}
		s.Cancel()
	}()
	<-connCtx.Done()
}

// ShowAllSessions 返回所有的session情况，对其他包提供，暂时先保留
func ShowAllSessions() {

}
