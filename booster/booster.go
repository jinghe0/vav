package booster

import (
	"github.com/giskook/vav/conf"
	tcp_server "github.com/giskook/vav/socket_server/tcp"
)

type Booster struct {
	conf *conf.Conf
	exit chan struct{}
	ss   *tcp_server.SocketServer
}

func NewBooster(conf *conf.Conf) *Booster {
	return &Booster{
		conf: conf,
		exit: make(chan struct{}),
		ss:   tcp_server.NewSocketServer(conf),
	}
}

func (b *Booster) Start() error {
	err := b.ss.Start()
	if err != nil {
		return err
	}

	b.do()

	return nil
}

func (b *Booster) Stop() {
	b.ss.Stop()
	close(b.exit)
}
