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

func (s *SocketServer) get_play_type(sim, channel string, info *vcbase.VavmsInfo) (int, string, string, error) {
	var data_type, conn_play_type string
	play_type := 0

	if info.Status == rc.STATUS_LIVE {
		data_type = info.LiveType
		conn_play_type = ss.CONN_PLAY_LIVE
	} else if info.Status == rc.STATUS_BACK {
		data_type = info.PlayBackType
		conn_play_type = ss.CONN_PLAY_BACK
	} else {
		err := errors.New(fmt.Sprintf("sim %s chan %s redis [status] shuld set to [%s] or [%s] now is [%s]", sim, channel, rc.STATUS_LIVE, rc.STATUS_BACK, info.Status))
		return 0, "", "", err
	}

	if data_type == rc.DATA_TYPE_AUDIO_VIDEO ||
		data_type == rc.DATA_TYPE_VIDEO {
		play_type &= 1
	}
	if data_type == rc.DATA_TYPE_TWO_WAY_INTERCOM ||
		data_type == rc.DATA_TYPE_LISTEN {
		play_type &= 2
	}

	return play_type, conn_play_type, data_type, nil
}

func (s *SocketServer) OnPrepare(c *ss.Connection, id, channel string) error {
	vavms_info, err := rc.GetInstance().GetVavmsInfo(id, redis_cli.GetIDChannelKey(id, channel), s.conf.UUID, redis_cli.VAVMS_STREAM_MEDIA)
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
	open_pipe := func(play_type, aname, vname string) (string, string, error) {
		var pipe_a, pipe_v string
		if play_type == rc.DATA_TYPE_AUDIO_VIDEO ||
			play_type == rc.DATA_TYPE_VIDEO {
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

		if play_type == rc.DATA_TYPE_AUDIO_VIDEO ||
			play_type == rc.DATA_TYPE_TWO_WAY_INTERCOM ||
			play_type == rc.DATA_TYPE_LISTEN ||
			play_type == rc.DATA_TYPE_BROADCAST {
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

	play_type, conn_play_type, data_type, err := s.get_play_type(id, channel, vavms_info)
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}

	ffmpeg_path, err := ffmpeg_symbol("vffmpeg_" + id + "_" + channel + conn_play_type)
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}
	path_a, path_v, err := open_pipe(data_type, "a"+conn_play_type, "v"+conn_play_type)
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}
	var cmd string
	switch play_type {
	case 3:
		cmd = fmt.Sprintf(s.conf.WorkSpace.FfmpegArgsAV, ffmpeg_path, vavms_info.Vcodec, path_v, vavms_info.Acodec, path_a, vavms_info.DomainInner+"/"+redis_cli.GetIDChannelKey(id, channel))
		break
	case 1:
		cmd = fmt.Sprintf(s.conf.WorkSpace.FfmpegArgsV, ffmpeg_path, vavms_info.Vcodec, path_v, vavms_info.DomainInner+"/"+redis_cli.GetIDChannelKey(id, channel))
	case 2:
		cmd = fmt.Sprintf(s.conf.WorkSpace.FfmpegArgsA, ffmpeg_path, vavms_info.Acodec, path_a, vavms_info.DomainInner+"/"+redis_cli.GetIDChannelKey(id, channel))
	}

	c.SetProperty(id, channel, conn_play_type, cmd)
	err = redis_cli.SetVehicleChanStatus(id, channel, rc.STATUS_INIT)
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}
	err = redis_cli.SetVehicleUUID(id, channel, s.conf.UUID)
	if err != nil {
		mybase.ErrorCheckPlus(err, id, channel)
		return err
	}

	return nil
}

func (s *SocketServer) OnClose(conn *ss.Connection) error {
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
