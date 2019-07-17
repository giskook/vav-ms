package main

import (
	"log"
	"os/exec"
)

func main() {
	//cmd := exec.Command("bash", "-c", "test/15226563111/3/vffmpeg_15226563111_3_live -y -thread_queue_size 128 -use_wallclock_as_timestamps 1 -f h264 -i test/15226563111/3/v_live -vcodec copy -f flv rtmp://127.0.0.1:8080/vavms/15226563111_3")
	cmd := exec.Command("bash", "-c", "ffmpeg -re -stream_loop -1 -i ./echo-hereweare.mp4  -c copy -f flv rtmp://127.0.0.1:8080/myapp/1")
	_, err := cmd.Output()
	if err != nil {
		log.Printf("<INFO> run ffmpeg error %s \n", err.Error())
	}
}
