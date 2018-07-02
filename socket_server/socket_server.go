package socket_server

import (
	"fmt"
	"github.com/giskook/vav/conf"
	"log"
	"net"
)

type SocketServer struct {
	cnf *conf.Conf
}

func NewSocketServer(cnf *conf.Conf) *SocketServer {
	return &SocketServer{
		cnf: cnf,
	}
}

func (s *SocketServer) Start() error {
	ServerAddr, err := net.ResolveUDPAddr("udp", s.cnf.UDP.Addr)
	if err != nil {
		log.Printf("<ERR> SocketServer %s\n", err.Error())
		return err
	}

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	defer ServerConn.Close()
	if err != nil {
		log.Println("<ERR> SocketServer %s\n", err.Error())
		return err
	}
	log.Printf("<INF> udp listening %s\n", s.cnf.UDP.Addr)

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	return nil
}
