package conf

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	t.Log(GetInstance().UUID)
	t.Log(*GetInstance().TCP)
	t.Log(*GetInstance().Redis)
	t.Log(*GetInstance().Http)
	t.Log(GetInstance().Formats)
	t.Log(GetInstance().SamplingRates)
	t.Log(GetInstance().CheckFormat(133))
	t.Log(GetInstance().CheckFormat(99))
	t.Log(GetInstance().CheckFormat(3))
	t.Log(GetInstance().CheckSamplingRate(3))
	t.Log(GetInstance().CheckSamplingRate(4))
}
