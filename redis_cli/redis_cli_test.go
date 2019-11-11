package redis_cli

import (
	"github.com/giskook/vav-ms/conf"
	"testing"
)

func TestInit(t *testing.T) {
	Init(conf.GetInstance())
}

//func TestPlayLock(t *testing.T) {
//	Init(conf.GetInstance())
//	t.Log(SetPlayLock("test_ttl", "live", "100"))
//}

func TestSetPlayInit(t *testing.T) {
	Init(conf.GetInstance())
	rt, err := StreamPlayInit(GetIDChannel("13731143001", "100", "live"), VAVMS_STREAM_DATA_TYPE_KEY, "0", VAVMS_STREAM_TTL_KEY, "1000", STREAM_PRIORITY_KEY, "10")
	t.Log(rt)
	t.Log(err)
}

func TestSetMaxPriority(t *testing.T) {
	Init(conf.GetInstance())
	StreamPlayInit(GetIDChannel("13731143001", "101", "live"), VAVMS_STREAM_DATA_TYPE_KEY, "0", VAVMS_STREAM_TTL_KEY, "1000", STREAM_PRIORITY_KEY, "10")
	_, err := StreamReplacePriority(GetIDChannel("13731143001", "101", "live"), STREAM_PRIORITY_KEY, "100")
	t.Log(err)
}
