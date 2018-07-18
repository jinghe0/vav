package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// catchs system signal
	cmd := exec.Command("/usr/local/bin/ffmpeg", "-re", "-i", "./echo-hereweare.mp4", "-vcodec", "copy", "-acodec", "copy", "-b:v", "800k", "-b:a", "32k", "-f", "flv", "rtmp://127.0.0.1:8080/myapp")
	err := cmd.Run()
	checkError(err)
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
