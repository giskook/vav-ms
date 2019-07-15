package socket_server

import (
	_ "github.com/giskook/vav-common/base"
	rc "github.com/giskook/vav-common/redis_cli"
	ss "github.com/giskook/vav-common/socket_server/tcp"
	"github.com/giskook/vav-ms/conf"
	"github.com/giskook/vav-ms/redis_cli"
	"log"
)

type SocketServer struct {
	server *ss.SocketServer
	conf   *conf.Conf
}

func (s *SocketServer) prepare(id, channel string) error {
	vavms_info, err := rc.GetInstance().GetVavmsInfo(id, redis_cli.GetIDChannelKey(id, channel), s.conf.UUID, redis_cli.AV_STREAM_MEDIA)
	if err != nil {
		return err
	}
	if vavms_info.PlayType == redis_cli.PLAY_TYPE_AV {

	}

	return err
}

func (s *SocketServer) OnDataAudio(conn *ss.Connection) bool {
	return true
}

func (s *SocketServer) OnDataVideo(conn *ss.Connection) bool {
	return true
}

func (s *SocketServer) OnClose(conn *ss.Connection) bool {
	return true
}

func NewSocketServer(conf *conf.Conf) *SocketServer {
	server := &SocketServer{
		conf: conf,
	}
	cnf := &ss.Conf{
		TcpAddr:         conf.TCP.Addr,
		ServerType:      ss.SERVER_TYPE_VAVMS,
		DefautReadLimit: conf.TCP.DefautReadLimit,
	}

	server.server = ss.NewSocketServer(cnf, server, server.prepare)

	return server
}

func (s *SocketServer) Start() error {
	return s.server.Start()
}

func (s *SocketServer) Stop() {
	s.server.Stop()
}
