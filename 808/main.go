package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cnf, err := ReadConfig("./conf.json")
	checkError(err)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", cnf.TcpSrvAddr)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	log.Printf("<INF> Listenning %s\n", cnf.TcpSrvAddr)
	for {
		conn, err := listener.AcceptTCP()
		checkError(err)

		go func() {
			data := make([]byte, 1024)
			length, err := conn.Read(data)
			checkError(err)
			if length == 0 {
				log.Println("<ERR> peer close")
				return
			}
			log.Printf("<INF> %x \n", data[0:length])
		}()
	}

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
