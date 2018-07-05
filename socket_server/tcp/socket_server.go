package socket_server

import (
	"github.com/gansidui/gotcp"
	"github.com/giskook/vav/base"
	"github.com/giskook/vav/conf"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

type SocketServer struct {
	conf      *conf.Conf
	srv       *gotcp.Server
	cm        *ConnMgr
	ChanFrame chan *base.Frame
	exit      chan struct{}
	wait_exit *sync.WaitGroup
	conn_uuid uint32
}

func NewSocketServer(conf *conf.Conf) *SocketServer {
	return &SocketServer{
		conf:      conf,
		cm:        NewConnMgr(),
		ChanFrame: make(chan *base.Frame),
		exit:      make(chan struct{}),
		wait_exit: new(sync.WaitGroup),
	}
}

func (ss *SocketServer) Start() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ss.conf.TCP.Addr)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	config := &gotcp.Config{
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}

	ss.srv = gotcp.NewServer(config, ss, ss)

	go ss.srv.Start(listener, time.Second)
	log.Println("<INFO> socket listening:", listener.Addr())

	return nil
}

func (ss *SocketServer) Stop() {
	close(ss.exit)
	ss.wait_exit.Wait()
	close(ss.ChanFrame)

	ss.srv.Stop()
}

func (ss *SocketServer) SetStatus(imei, status string) *Connection {
	id, _ := strconv.ParseUint(imei, 10, 64)
	s, _ := strconv.ParseUint(status, 10, 8)
	c := ss.cm.Get(id)
	if c != nil {
		c.status = uint8(s)
	}

	return c
}
