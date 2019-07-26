package redis_cli

import (
	"github.com/giskook/vav-ms/conf"
	"testing"
)

func TestInit(t *testing.T) {
	Init(conf.GetInstance())
}

func TestPlayLock(t *testing.T) {
	Init(conf.GetInstance())
	t.Log(SetPlayLock("test_ttl", "live", "100"))
}
