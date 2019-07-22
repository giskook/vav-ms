package v1

import (
	"encoding/json"
	gkbase "github.com/giskook/go/base"
	gkhttp "github.com/giskook/go/http"
	vcbase "github.com/giskook/vav-common/base"
	//rc "github.com/giskook/vav-common/redis_cli"
	"github.com/giskook/vav-ms/base"
	"github.com/giskook/vav-ms/redis_cli"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

const (
	STREAM_TYPE_LIVE string = "live"
	STREAM_TYPE_BACK string = "back"
)

type stream struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

func stream_condition(sim string) (error, error, error) {
	var wg sync.WaitGroup
	var ip, port, audio_format, video_format string
	var err1, err2, err3 error
	var stream_media []*vcbase.StreamMedia
	go func() {
		wg.Add(1)
		ip, port, err1 = redis_cli.GetAccessAddr()
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		audio_format, video_format, err2 = redis_cli.GetVehicleProperty(sim)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		stream_media, err3 = redis_cli.GetStreamMedia()
		wg.Done()
	}()
	wg.Wait()

	return err1, err2, err3
}

func stream_post(w http.ResponseWriter, r *http.Request) (int, string, error) {
	r.ParseForm()
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var s stream
	err := decoder.Decode(&s)
	if err != nil {
		gkbase.ErrorCheck(err)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_DECODE, err
	}
	vars := mux.Vars(r)
	sim := vars["sim"]
	//channel := vars["channel"]
	if (s.Type == "") ||
		(s.Type != STREAM_TYPE_LIVE && s.Type != STREAM_TYPE_BACK) {
		gkbase.ErrorCheck(base.ERROR_BAD_REQUEST_MISSING)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_MISSING, base.ERROR_BAD_REQUEST_MISSING
	}

	err1, err2, err3 := stream_condition(sim)
	if err1 != nil {
		return http.StatusAccepted, base.HTTP_ACCEPTED_AV_ACCESS_ADDR_NOT_SET, nil
	}
	if err2 != nil {
		return http.StatusAccepted, base.HTTP_ACCEPTED_AV_FORMAT_NOT_SET, nil
	}
	if err3 != nil {
		return http.StatusAccepted, base.HTTP_ACCEPTED_AV_STREAM_MEDIA_NOT_SET, nil
	}

	if s.Type == STREAM_TYPE_LIVE {
	}

	return http.StatusCreated, base.HTTP_OK, nil
}

func Stream(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			gkbase.ErrorPrintStack()
			w.WriteHeader(http.StatusInternalServerError)
			gkhttp.EncodeResponse(w, base.HTTP_INTERNAL_SERVER_ERROR, nil, "")
		}
	}()
	gkhttp.RecordReq(r)

	var http_status int
	var internal_status string
	var err error
	var data interface{}

	switch r.Method {
	case http.MethodPost:
		http_status, internal_status, err = stream_media_post(w, r)
	}

	common_reply(w, http_status, internal_status, data, err)
}
