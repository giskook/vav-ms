package main

import (
	"fmt"
	"github.com/giskook/vav-ms/conf"
	"github.com/giskook/vav-ms/redis_cli"
	vavms "github.com/giskook/vav-ms/socket_server"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	redis_cli.Init(conf.GetInstance())
	server := vavms.NewSocketServer(conf.GetInstance())
	err := server.Start()
	if err != nil {
		log.Fatal(err.Error())
	}

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
	server.Stop()
}
