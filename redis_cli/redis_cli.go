package redis_cli

import (
	vcbase "github.com/giskook/vav-common/base"
	rc "github.com/giskook/vav-common/redis_cli"
	"github.com/giskook/vav-ms/conf"
	"strings"
)

const (
	VAVMS_STREAM_MEDIA         string = "vavms_stream_media"
	VAVMS_ACCESS_ADDR          string = "vavms_access_addr"
	VAVMS_STREAM_DATA_TYPE_KEY string = "data_type"
	VAVMS_STREAM_TTL_KEY       string = "ttl"
	VAVMS_STREAM_URL_KEY       string = "url"

	VAVMS_PLAY_STATUS_LIVE string = "live"
	VAVMS_PLAY_STATUS_BACK string = "back"

	VAVMS_ACCESS_ADDR_UUID string = "uuid"
)

func GetIDChannel(args ...string) string {
	return strings.Join(args, "_")
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

func SetAccessAddr(ip, port string) error {
	return rc.GetInstance().SetAccessAddr(VAVMS_ACCESS_ADDR, ip, port)
}

func GetAccessAddr() (string, string, error) {
	return rc.GetInstance().GetAccessAddr(VAVMS_ACCESS_ADDR)
}

func SetVehicleProperty(sim, audio_format, video_format, sampling_rate string) error {
	return rc.GetInstance().SetVehicleStreamFormat(sim, audio_format, video_format, sampling_rate)
}

func GetVehicleProperty(sim string) (string, string, error) {
	return rc.GetInstance().GetVehicleStreamFormat(sim)
}

func GetStreamMedia() ([]*vcbase.StreamMedia, error) {
	return rc.GetInstance().GetStreamMedia(VAVMS_STREAM_MEDIA, "0", "-1")
}

func SetStreamMedia(s []*vcbase.StreamMedia) error {
	return rc.GetInstance().SetStreamMedia(VAVMS_STREAM_MEDIA, s)
}

func DelStreamMedia(index string) bool {
	return rc.GetInstance().DelStreamMedia(VAVMS_STREAM_MEDIA, index)
}

func UpdateStreamMedia(index string, new_stream_dedia *vcbase.StreamMedia) bool {
	return rc.GetInstance().UpdateStreamMedia(VAVMS_STREAM_MEDIA, index, new_stream_dedia)
}

func SetVehicleUUID(id, channel, uuid string) error {
	return rc.GetInstance().SetVehicleChan(GetIDChannel(id, channel), VAVMS_ACCESS_ADDR_UUID, uuid)
}

func SetPlayLock(key, status, ttl string) (int, error) {
	return rc.GetInstance().SetPlayLock(key, status, ttl)
}

func AddKey(key, value string) error {
	return rc.GetInstance().AddKey(key, value)
}
