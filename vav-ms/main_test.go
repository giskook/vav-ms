package main

import (
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	go func() {
		time.Sleep(5 * time.Second)
		f, err := os.OpenFile("test/15210660145/1/vlive", os.O_RDWR, 0600)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("close ")
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("run ffmpeg")
	cmd := exec.Command("bash", "-c", "test/15210660145/1/vffmpeg_15210660145_1_live -y -thread_queue_size 128 -use_wallclock_as_timestamps 1 -f h264 -i test/15210660145/1/vlive -vcodec copy -f flv rtmp://127.0.0.1:8080/myapp/15210660145_1_live")
	log.Println("run ffmpeg")
	cmd.Output()
	log.Println("ffmpeg exit")

}
