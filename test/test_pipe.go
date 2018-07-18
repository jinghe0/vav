package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// catchs system signal
	go func() {
		log.Println("open r")
		//f, e := os.Open("/tmp/testpipe")
		f, e := os.OpenFile("/tmp/testpipe", os.O_RDONLY, 0600)
		log.Println("open r")
		checkError(e)
		for {
			b := make([]byte, 1024)
			f.Read(b)
			log.Printf("%x\n", b)
			//f.Write([]byte("hello world"))
			//time.Sleep(1 * time.Second)
			//log.Println("hi")
		}
	}()
	//	go func() {
	//		log.Println("open w")
	//		f, e := os.OpenFile("/tmp/testpipe", os.O_RDWR, 0600)
	//		log.Println("open w")
	//		checkError(e)
	//		for {
	//			b := make([]byte, 1024)
	//			f.Read(b)
	//			log.Println(string(b))
	//		}
	//	}()
	//r, w, e := os.Pipe()
	//log.Println(r)
	//log.Println(w)
	//checkError(e)
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
