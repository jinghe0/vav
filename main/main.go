package main

import (
	"fmt"
	"github.com/giskook/vav/conf"
	"github.com/giskook/vav/socket_server"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	conf, err := conf.ReadConfig("./conf.json")
	checkError(err)
	ss := socket_server.NewSocketServer(conf)
	ss.Start()
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
