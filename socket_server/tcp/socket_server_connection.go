package socket_server

import (
	"bytes"
	"github.com/gansidui/gotcp"
	"time"
)

const (
	USER_STATUS_INIT    uint8 = 0
	USER_STATUS_NORMAL  uint8 = 1
	USER_STATUS_ILLEGAL uint8 = 2
)

type ConnConf struct {
	read_limit int
	uuid       uint32
}

type Connection struct {
	conf       *ConnConf
	c          *gotcp.Conn
	ID         uint64
	RecvBuffer *bytes.Buffer
	exit       chan struct{}
	status     uint8
	Buffer     []byte
}

func NewConnection(c *gotcp.Conn, conf *ConnConf) *Connection {
	tcp_c := c.GetRawConn()
	tcp_c.SetReadDeadline(time.Now().Add(time.Duration(conf.read_limit) * time.Second))
	return &Connection{
		conf:       conf,
		c:          c,
		RecvBuffer: bytes.NewBuffer([]byte{}),
		exit:       make(chan struct{}),
	}
}

func (c *Connection) SetReadDeadline() {
	c.c.GetRawConn().SetReadDeadline(time.Now().Add(time.Duration(c.conf.read_limit) * time.Second))
}

func (c *Connection) Close() {
	close(c.exit)
	c.RecvBuffer.Reset()
}

func (c *Connection) Equal(cc *Connection) bool {
	return c.conf.uuid == cc.conf.uuid
}
