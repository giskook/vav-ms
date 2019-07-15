package socket_server

import (
	"github.com/giskook/vav-ms/conf"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestNewSocketServer(t *testing.T) {
	server := NewSocketServer(conf.GetInstance())
	err := server.Start()
	if err != nil {
		t.Fatal(err)
	}
	// signal not work
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	server.Stop()
}
