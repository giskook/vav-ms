package socket_server

import (
	"errors"
	"fmt"
	mybase "github.com/giskook/go/base"
	vcbase "github.com/giskook/vav-common/base"
	rc "github.com/giskook/vav-common/redis_cli"
	ss "github.com/giskook/vav-common/socket_server/tcp"
	"github.com/giskook/vav-ms/conf"
	"github.com/giskook/vav-ms/redis_cli"
	"log"
	"path"
	"strconv"
	"time"
)

const (
	RTMP_FORMATER string = "rtmp://%s:%s/%s/%s"
)

type SocketServer struct {
	server *ss.SocketServer
	conf   *conf.Conf
}

func (s *SocketServer) OnFfmpegExit(conn *ss.Connection) error {
	//	redis_cli.StreamDelUrl(redis_cli.GetIDChannel(conn.SIM, conn.Channel, conn.PlayType))
	//	conn.Close()

	return nil
}

//func (s *SocketServer) get_play_type(info *vcbase.VavmsInfo) int {
//	play_type := 0
//	if info.DataType == rc.DATA_TYPE_AUDIO_VIDEO {
//		play_type |= 3
//	}
//	if info.DataType == rc.DATA_TYPE_VIDEO {
//		play_type |= 1
//	}
//	if info.DataType == rc.DATA_TYPE_TWO_WAY_INTERCOM ||
//		info.DataType == rc.DATA_TYPE_LISTEN {
//		play_type |= 2
//	}
//
//	return play_type
//}

func (s *SocketServer) OnPrepare(c *ss.Connection, id, channel string) error {
	vavms_info, err := rc.GetInstance().GetVavmsInfo(id, channel, s.conf.UUID, redis_cli.VAVMS_STREAM_MEDIA)
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}
	if vavms_info.Acodec == "" ||
		vavms_info.Vcodec == "" ||
		vavms_info.StreamMedia == nil {
		e := errors.New("redis not configured")
		log.Println(vavms_info)
		mybase.ErrorCheckPlus(e, id, channel)
		return e
	}

	// pipe path and open pipe
	open_pipe := func(aname, vname string) (string, string, error) {
		var pipe_a, pipe_v string
		pipe_v = path.Join(s.conf.WorkSpace.PipePath, id, channel, vname)
		err = vcbase.Mkfifo(pipe_v)
		if err != nil {
			mybase.ErrorCheckPlus(err, id, channel)
			return "", "", err
		}
		err = c.OpenPipeV(pipe_v)
		if err != nil {
			mybase.ErrorCheckPlus(err, id, channel)
			return "", "", err
		}

		pipe_a = path.Join(s.conf.WorkSpace.PipePath, id, channel, aname)
		err = vcbase.Mkfifo(pipe_a)
		if err != nil {
			mybase.ErrorCheckPlus(err, id, channel)
			return "", "", err
		}
		err = c.OpenPipeA(pipe_a)
		if err != nil {
			mybase.ErrorCheckPlus(err, id, channel)
			return "", "", err
		}

		return pipe_a, pipe_v, nil
	}
	// create ffmpeg and start
	ffmpeg_symbol := func(name string) (string, error) {
		// create ffmpeg symbol linke vavms shows this project
		ffmpeg_path := path.Join(s.conf.WorkSpace.PipePath, id, channel, name)
		err = vcbase.Symlink(s.conf.WorkSpace.FfmpegBin, ffmpeg_path)
		if err != nil {
			mybase.ErrorCheckPlus(err, id, channel)
			return "", err
		}

		return ffmpeg_path, nil
	}

	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}

	ttl, err := strconv.Atoi(vavms_info.TTL)
	if err != nil {
		ttl = 500
	}
	tmp_t := time.Now().Add(time.Duration(ttl) * time.Second).Unix()

	ffmpeg_name := redis_cli.GetIDChannel("ffmpeg", id, channel, vavms_info.Status)
	ffmpeg_path, err := ffmpeg_symbol(ffmpeg_name + "_" + strconv.Itoa(int(tmp_t)) + "_" + vavms_info.TTL)
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}
	path_a, path_v, err := open_pipe("a"+vavms_info.Status, "v"+vavms_info.Status)
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}
	cmds := make([]string, 0)
	url_inner := fmt.Sprintf(RTMP_FORMATER, vavms_info.StreamMedia.RtmpIpInner, vavms_info.StreamMedia.RtmpPortInner, vavms_info.StreamMedia.RtmpApplication, redis_cli.GetIDChannel(id, channel, vavms_info.Status))
	url_outer := fmt.Sprintf(RTMP_FORMATER, vavms_info.StreamMedia.RtmpIpOutter, vavms_info.StreamMedia.RtmpPortOutter, vavms_info.StreamMedia.RtmpApplication, redis_cli.GetIDChannel(id, channel, vavms_info.Status))
	cmds = append(cmds, fmt.Sprintf(s.conf.WorkSpace.FfmpegArgsV, ffmpeg_path, vavms_info.Vcodec, path_v, url_inner))
	cmds = append(cmds, fmt.Sprintf(s.conf.WorkSpace.FfmpegArgsA, ffmpeg_path, vavms_info.Acodec, vavms_info.SamplingRate, path_a, url_inner))
	cmds = append(cmds, fmt.Sprintf(s.conf.WorkSpace.FfmpegArgsAV, ffmpeg_path, vavms_info.Vcodec, path_v, vavms_info.Acodec, vavms_info.SamplingRate, path_a, url_inner))

	result, err := redis_cli.StreamDestruct(redis_cli.GetIDChannel(id, channel, vavms_info.Status),
		redis_cli.VAVMS_ACCESS_ADDR_UUID, s.conf.UUID,
		redis_cli.VAVMS_STREAM_URL_KEY, url_outer,
		rc.HTTP_IP_OUTTER, vavms_info.StreamMedia.HttpIpOutter,
		rc.HTTP_PORT_OUTTER, vavms_info.StreamMedia.HttpPortOutter,
		rc.HTTP_LOCATION, vavms_info.StreamMedia.HttpLocation,
		rc.RTMP_INNER_PORT, vavms_info.StreamMedia.RtmpPortInner,
		rc.RTMP_APPLICATION, vavms_info.StreamMedia.RtmpApplication,
		redis_cli.VAVMS_STREAM_TTL_KEY,
		redis_cli.GetIDChannel(id, channel, "status"))
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel, vavms_info.Status)
		return err
	}
	if result != 0 {
		e := errors.New(fmt.Sprintf("prepare destruct sim %s chan %s result %d", id, channel, result))
		mybase.ErrorCheckPlus(e, id, channel, vavms_info.Status)
		return e
	}

	c.SetProperty(id, channel, vavms_info.Status, path_a, path_v, vavms_info.Acodec, vavms_info.Vcodec, ffmpeg_name, cmds, ttl)

	return nil
}

func (s *SocketServer) OnClose(conn *ss.Connection) error {
	err := redis_cli.StreamDelUrl(redis_cli.GetIDChannel(conn.SIM, conn.Channel, conn.PlayType))
	if err != nil {
		mybase.ErrorCheckPlus(err, conn.SIM, conn.Channel, conn.PlayType)
	}
	return nil
}

func NewSocketServer(conf *conf.Conf) *SocketServer {
	server := &SocketServer{
		conf: conf,
	}
	cnf := &ss.Conf{
		TcpAddr:          conf.TCP.Addr,
		ServerType:       ss.SERVER_TYPE_VAVMS,
		DefaultReadLimit: conf.TCP.DefaultReadLimit,
		FFmpegKiller:     conf.WorkSpace.FfmpegKiller,
		Debug: &ss.DebugCnf{
			Debug:       conf.Debug.Debug,
			DestID:      conf.Debug.DestID,
			RecordFileA: conf.Debug.RecordFileA,
		},
	}

	server.server = ss.NewSocketServer(cnf, server)

	return server
}

func (s *SocketServer) Start() error {
	return s.server.Start()
}

func (s *SocketServer) Stop() {
	s.server.Stop()
}
