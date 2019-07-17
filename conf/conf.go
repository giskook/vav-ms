package conf

import (
	"encoding/json"
	mybase "github.com/giskook/go/base"
	"log"
	"os"
	"sync"
	"time"
)

const (
	CONF_FILE string = "./conf.json"
)

type tcp_srv_cnf struct {
	Addr              string `json:"addr"`
	DefaultReadLimit  time.Duration
	TDefaultReadLimit string `json:"t_default_read_limit"`
}

type redis_cnf struct {
	Addr          string        `json:"addr"`
	Passwd        string        `json:"passwd"`
	MaxIdle       int           `json:"max_idle"`
	TConnTimeOut  string        `json:"t_conn_time_out"`
	ConnTimeOut   time.Duration `json:"conn_time_out"`
	TReadTimeOut  string        `json:"t_read_time_out"`
	ReadTimeOut   time.Duration `json:"read_time_out"`
	TWriteTimeOut string        `json:"t_write_time_out"`
	WriteTimeOut  time.Duration `json:"write_time_out"`
}

type workspace_cnf struct {
	PipePath     string `json:"pipe_path"`
	FfmpegBin    string `json:"ffmpeg_bin"`
	FfmpegArgsAV string `json:"ffmpeg_args_av"`
	FfmpegArgsA  string `json:"ffmpeg_args_a"`
	FfmpegArgsV  string `json:"ffmpeg_args_v"`
}

type Conf struct {
	UUID      string         `json:"uuid"`
	TCP       *tcp_srv_cnf   `json:"tcp"`
	Redis     *redis_cnf     `json:"redis"`
	WorkSpace *workspace_cnf `json:"work_space"`
}

var instance Conf
var once sync.Once

func GetInstance() *Conf {
	once.Do(func() {
		file, _ := os.Open(CONF_FILE)
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&instance)
		error_check(err)
		instance.TCP.DefaultReadLimit, err = time.ParseDuration(instance.TCP.TDefaultReadLimit)
		error_check(err)
		instance.Redis.ConnTimeOut, err = time.ParseDuration(instance.Redis.TConnTimeOut)
		error_check(err)
		instance.Redis.ReadTimeOut, err = time.ParseDuration(instance.Redis.TReadTimeOut)
		error_check(err)
		instance.Redis.WriteTimeOut, err = time.ParseDuration(instance.Redis.TWriteTimeOut)
		error_check(err)

	})

	return &instance
}

func error_check(err error) {
	if err != nil {
		mybase.ErrorCheckWithLevel(err, mybase.UPPER_LEVEL)
		log.Fatal(err.Error())
	}
}
