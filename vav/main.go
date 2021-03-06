package main

import (
	"fmt"
	"github.com/giskook/vav/booster"
	"github.com/giskook/vav/conf"
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
	booster := booster.NewBooster(conf)
	booster.Start()
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
