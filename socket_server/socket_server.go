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
)

type SocketServer struct {
	server *ss.SocketServer
	conf   *conf.Conf
}

func (s *SocketServer) OnPrepare(c *ss.Connection, id, channel string) error {
	vavms_info, err := rc.GetInstance().GetVavmsInfo(id, redis_cli.GetIDChannelKey(id, channel), s.conf.UUID, redis_cli.AV_STREAM_MEDIA)
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}
	if vavms_info.Acodec == "" ||
		vavms_info.Vcodec == "" ||
		vavms_info.DomainInner == "" {
		e := errors.New("redis not configured")
		log.Println(vavms_info)
		mybase.ErrorCheckPlus(e, id, channel)
		return e
	}

	// pipe path and open pipe
	open_pipe := func(play_type int, aname, vname string) (string, string, error) {
		var pipe_a, pipe_v string
		if (play_type & rc.PLAY_TYPE_V) != 0 {
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
		}

		if (play_type & rc.PLAY_TYPE_A) != 0 {
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

	var suffix string
	var play_type int
	if vavms_info.LiveStatus == rc.PLAY_STATUS_INIT && vavms_info.PlayBackStatus != rc.PLAY_STATUS_INIT { // live
		suffix = "_live"
		play_type = vavms_info.LiveType
	} else if vavms_info.LiveStatus != rc.PLAY_STATUS_INIT && vavms_info.PlayBackStatus == rc.PLAY_STATUS_INIT { // play back
		suffix = "_playback"
		play_type = vavms_info.PlayBackType
	} else {
		e := errors.New(fmt.Sprintf("redis not ok live status %d playback status %d ", id, channel, vavms_info.LiveStatus, vavms_info.PlayBackStatus))
		mybase.ErrorCheckPlus(e, id, channel)
		return e
	}

	ffmpeg_path, err := ffmpeg_symbol("vffmpeg_" + id + "_" + channel + suffix)
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}
	path_a, path_v, err := open_pipe(play_type, "a"+suffix, "v"+suffix)
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}
	var cmd string
	switch vavms_info.LiveType {
	case 3: // both a and v
		cmd = fmt.Sprintf(s.conf.WorkSpace.FfmpegArgsAV, ffmpeg_path, vavms_info.Vcodec, path_v, vavms_info.Acodec, path_a, vavms_info.DomainInner+"/"+redis_cli.GetIDChannelKey(id, channel))
		break
	case 1: // v
		cmd = fmt.Sprintf(s.conf.WorkSpace.FfmpegArgsV, ffmpeg_path, vavms_info.Vcodec, path_v, vavms_info.DomainInner+"/"+redis_cli.GetIDChannelKey(id, channel))
		break
	case 2: // a
		cmd = fmt.Sprintf(s.conf.WorkSpace.FfmpegArgsA, ffmpeg_path, vavms_info.Acodec, path_a, vavms_info.DomainInner+"/"+redis_cli.GetIDChannelKey(id, channel))
		break
	}

	c.SetFfmpegCmd(cmd)

	return nil
}

func (s *SocketServer) OnClose(conn *ss.Connection) bool {
	return true
}

func NewSocketServer(conf *conf.Conf) *SocketServer {
	server := &SocketServer{
		conf: conf,
	}
	cnf := &ss.Conf{
		TcpAddr:          conf.TCP.Addr,
		ServerType:       ss.SERVER_TYPE_VAVMS,
		DefaultReadLimit: conf.TCP.DefaultReadLimit,
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
