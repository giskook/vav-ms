package v1

import (
	"encoding/json"
	gkbase "github.com/giskook/go/base"
	gkhttp "github.com/giskook/go/http"
	"github.com/giskook/vav-ms/base"
	"github.com/giskook/vav-ms/conf"
	"github.com/giskook/vav-ms/redis_cli"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type vehicle_property struct {
	AudioFormat       string `json:"audio_format"`
	AudioSamplingRate string `json:"audio_sampling_rate"`
	VideoFormat       string `json:"video_format"`
}

func vehicle_property_post(w http.ResponseWriter, r *http.Request) (int, string, error) {
	vars := mux.Vars(r)
	sim := vars["sim"]
	r.ParseForm()
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var vp vehicle_property
	err := decoder.Decode(&vp)
	if err != nil {
		gkbase.ErrorCheck(err)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_DECODE, err
	}
	if vp.AudioFormat == "" ||
		vp.VideoFormat == "" ||
		vp.AudioSamplingRate == "" {
		gkbase.ErrorCheck(base.ERROR_BAD_REQUEST_MISSING)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_MISSING, base.ERROR_BAD_REQUEST_MISSING
	}

	// convert to ffmpeg supports formts
	audio_format_index, err1 := strconv.Atoi(vp.AudioFormat)
	video_format_index, err2 := strconv.Atoi(vp.VideoFormat)
	audio_sampling_rate_index, err3 := strconv.Atoi(vp.AudioSamplingRate)
	if err1 != nil || err2 != nil || err3 != nil {
		gkbase.ErrorCheck(base.ERROR_BAD_REQUEST_MISSING)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_MISSING, base.ERROR_BAD_REQUEST_MISSING
	}
	audio_format, err := conf.GetInstance().CheckFormat(audio_format_index)
	if err != nil {
		gkbase.ErrorCheck(base.ERROR_BAD_REQUEST_MISSING)
		if err == base.ERROR_BAD_REQUEST_ASSETS_OVER_RANGE {
			return http.StatusBadRequest, base.HTTP_BAD_REQUEST_ASSETS_OVER_RANGE, err
		}
		return http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR_FORMATS_NOT_SUPPORT, err
	}

	video_format, err := conf.GetInstance().CheckFormat(video_format_index)
	if err != nil {
		gkbase.ErrorCheck(base.ERROR_BAD_REQUEST_MISSING)
		if err == base.ERROR_BAD_REQUEST_ASSETS_OVER_RANGE {
			return http.StatusBadRequest, base.HTTP_BAD_REQUEST_ASSETS_OVER_RANGE, err
		}
		return http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR_FORMATS_NOT_SUPPORT, err
	}

	sampling_rate, err := conf.GetInstance().CheckSamplingRate(audio_sampling_rate_index)
	if err != nil {
		gkbase.ErrorCheck(base.ERROR_BAD_REQUEST_MISSING)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_ASSETS_OVER_RANGE, err
	}

	err = redis_cli.SetVehicleProperty(sim, audio_format, video_format, sampling_rate)
	if err != nil {
		gkbase.ErrorCheck(err)
		return http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR_SET_VEHICLE_PROPERTY, err
	}

	return http.StatusCreated, base.HTTP_OK, nil
}

func VehicleProperty(w http.ResponseWriter, r *http.Request) {
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
		http_status, internal_status, err = vehicle_property_post(w, r)
	}

	common_reply(w, http_status, internal_status, data, err)
}
