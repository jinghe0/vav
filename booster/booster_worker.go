package booster

import (
	"log"
	"os"
	//"os/exec"
	//"strconv"
)

func (b *Booster) booster_do() {
	go func() {
		//cmd := exec.Command("/usr/local/bin/ffmpeg", "-re", "-i", "./echo-hereweare.mp4", "-vcodec", "copy", "-acodec", "copy", "-b:v", "800k", "-b:a", "32k", "-f", "flv", "rtmp://127.0.0.1:8080/myapp")
		//// ffmpeg -f h264 -i <stream.h264>  -an -vcodec copy -f flv rtmp://<SERVERIP>
		//	err := cmd.Run()
	}()
	f, e := os.OpenFile("/tmp/testpipe", os.O_RDWR, 0600)
	if e != nil {
		log.Println(e.Error())
	}

	for {
		select {
		case <-b.exit:
			return
		case p := <-b.ss.ChanFrame:
			log.Printf("<INF> sim %s chan %s type %x, len %d timestamp %d lastI %d lastF %d \n", p.SIM, p.LogicalChannel, p.Type, len(p.Data), p.Timestamp, p.LastIFrameInterval, p.LastFrameInterval)
			f.Write(p.Data)
		}
	}
}

func (b *Booster) do() {
	for i := 0; i < b.conf.Booster.WorkerNum; i++ {
		go b.booster_do()
	}
}
