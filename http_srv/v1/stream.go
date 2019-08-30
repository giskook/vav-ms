package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	gkbase "github.com/giskook/go/base"
	gkhttp "github.com/giskook/go/http"
	vcbase "github.com/giskook/vav-common/base"
	"strconv"
	//rc "github.com/giskook/vav-common/redis_cli"
	"github.com/giskook/vav-ms/base"
	"github.com/giskook/vav-ms/conf"
	"github.com/giskook/vav-ms/redis_cli"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
	"time"
)

const (
	STREAM_TYPE_LIVE string = "live"
	STREAM_TYPE_BACK string = "back"
)

type stream struct {
	DataType string `json:"data_type"`
	TTL      string `json:"ttl"`
}

func stream_condition(sim string) (error, error, error) {
	var wg sync.WaitGroup
	var ip, port, audio_format, video_format string
	var err1, err2, err3 error
	var stream_media []*vcbase.StreamMedia
	wg.Add(3)
	go func() {
		ip, port, err1 = redis_cli.GetAccessAddr()
		if ip == "" || port == "" {
			err1 = errors.New("access addr ip or port is empty")
		}
		wg.Done()
	}()
	go func() {
		audio_format, video_format, err2 = redis_cli.GetVehicleProperty(sim)
		if audio_format == "" || video_format == "" {
			err2 = errors.New("vehcile audio format or video format is empty")
		}
		wg.Done()
	}()
	go func() {
		stream_media, err3 = redis_cli.GetStreamMedia()
		if len(stream_media) == 0 {
			err3 = errors.New("stream media is empty")
		}
		wg.Done()
	}()
	wg.Wait()

	return err1, err2, err3
}

func stream_get(w http.ResponseWriter, r *http.Request) (int, string, interface{}, error) {
	r.ParseForm()
	vars := mux.Vars(r)
	stream_type := vars["type"]
	sim := vars["sim"]
	channel := vars["channel"]

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.GetInstance().Play.PlayLockTTL)*time.Second)
	defer cancel()

	for {
		select {
		case <-time.After(1 * time.Second):
			url, _ := redis_cli.StreamPlayURL(redis_cli.GetIDChannel(sim, channel, stream_type))
			if url != "" {
				return http.StatusOK, base.HTTP_OK, url, nil
			}
		case <-ctx.Done():
			return http.StatusInternalServerError, base.HTTP_INTERNAL_STREAM_TIMEOUT, nil, nil

		}
	}

	//return http.StatusOK, base.HTTP_OK, url, nil
}

type stream_post_response struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func stream_post(w http.ResponseWriter, r *http.Request) (int, string, interface{}, error) {
	r.ParseForm()
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var stream stream
	err := decoder.Decode(&stream)
	if err != nil {
		gkbase.ErrorCheck(err)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_DECODE, nil, err
	}
	if stream.DataType == "" ||
		stream.TTL == "" {
		gkbase.ErrorCheck(base.ERROR_BAD_REQUEST_MISSING)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_MISSING, nil, base.ERROR_BAD_REQUEST_MISSING
	}

	if stream.DataType != "0" && stream.DataType != "1" && stream.DataType != "2" && stream.DataType != "3" && stream.DataType != "4" && stream.DataType != "5" {
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_FIELD_ILLEGAL, nil, err
	}

	vars := mux.Vars(r)
	stream_type := vars["type"]
	sim := vars["sim"]
	channel := vars["channel"]

	err1, err2, err3 := stream_condition(sim)
	if err1 != nil {
		return http.StatusAccepted, base.HTTP_ACCEPTED_AV_ACCESS_ADDR_NOT_SET, nil, nil
	}
	if err2 != nil {
		return http.StatusAccepted, base.HTTP_ACCEPTED_AV_FORMAT_NOT_SET, nil, nil
	}
	if err3 != nil {
		return http.StatusAccepted, base.HTTP_ACCEPTED_AV_STREAM_MEDIA_NOT_SET, nil, nil
	}

	result, err := redis_cli.SetPlayLock(redis_cli.GetIDChannel(sim, channel, "status"), stream_type, strconv.Itoa(conf.GetInstance().Play.PlayLockTTL))
	if err != nil {
		return http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR_UPD_STREAM_MEDIA, nil, nil
	}
	if result == 1 {
		return http.StatusConflict, base.HTTP_CONFLICT_PLAY, nil, nil
	}
	result, err = redis_cli.StreamPlayInit(redis_cli.GetIDChannel(sim, channel, stream_type), redis_cli.VAVMS_STREAM_DATA_TYPE_KEY, stream.DataType, redis_cli.VAVMS_STREAM_TTL_KEY, stream.TTL)
	if err != nil || result != 0 {
		gkbase.ErrorCheckPlus(err, fmt.Sprintf("stream type %s, sim %s, channel %s result %d", stream_type, sim, channel, result))
		return http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR, nil, nil

	}

	ip, port, _ := redis_cli.GetAccessAddr()

	return http.StatusCreated, base.HTTP_OK, &stream_post_response{
		IP:   ip,
		Port: port,
	}, nil
}

func Stream(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			gkbase.ErrorPrintStack()
			common_reply(w, http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR, nil, nil)
		}
	}()
	gkhttp.RecordReq(r)

	var http_status int
	var internal_status string
	var err error
	var data interface{}

	switch r.Method {
	case http.MethodGet:
		http_status, internal_status, data, err = stream_get(w, r)
	case http.MethodPost:
		http_status, internal_status, data, err = stream_post(w, r)
	}

	common_reply(w, http_status, internal_status, data, err)
}
