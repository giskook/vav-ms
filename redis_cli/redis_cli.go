package redis_cli

import (
	rc "github.com/giskook/vav-common/redis_cli"
	"github.com/giskook/vav-ms/conf"
)

const (
	AV_STREAM_MEDIA string = "vavms_stream_media"
)

func GetIDChannelKey(id, channel string) string {
	return id + "_" + channel
}

func Init(conf *conf.Conf) {
	cnf := &rc.Conf{
		Addr:         conf.Redis.Addr,
		Passwd:       conf.Redis.Passwd,
		MaxIdle:      conf.Redis.MaxIdle,
		ConnTimeOut:  conf.Redis.ConnTimeOut,
		ReadTimeOut:  conf.Redis.ReadTimeOut,
		WriteTimeOut: conf.Redis.WriteTimeOut,
	}
	rc.GetInstance().Init(cnf)
}

func SetStatus(id_channel, play_type, status string) error {
	return rc.GetInstance().SetVehicleChan(id_channel, play_type, status)
}
