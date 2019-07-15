package conf

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	t.Log(GetInstance().UUID)
	t.Log(*GetInstance().TCP)
	t.Log(*GetInstance().Redis)
}
