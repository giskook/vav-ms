package redis_cli

import (
	"github.com/giskook/vav-ms/conf"
	"testing"
)

func TestInit(t *testing.T) {
	Init(conf.GetInstance())
}
