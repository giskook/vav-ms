package v1

import (
	"encoding/json"
	gkbase "github.com/giskook/go/base"
	gkhttp "github.com/giskook/go/http"
	vcbase "github.com/giskook/vav-common/base"
	"github.com/giskook/vav-ms/base"
	"github.com/giskook/vav-ms/redis_cli"
	"github.com/gorilla/mux"
	"net/http"
)

func stream_media_get(w http.ResponseWriter, r *http.Request) (int, string, interface{}, error) {
	stream_medias, err := redis_cli.GetStreamMedia()
	if err != nil {
		return http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR_GET_STREAM_MEDIA, nil, err
	}

	return http.StatusOK, base.HTTP_OK, stream_medias, nil
}

func stream_media_post(w http.ResponseWriter, r *http.Request) (int, string, error) {
	r.ParseForm()
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var stream_medias base.StreamMedias
	err := decoder.Decode(&stream_medias)
	if err != nil {
		gkbase.ErrorCheck(err)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_DECODE, err
	}
	if stream_medias.StreamMedias == nil ||
		len(stream_medias.StreamMedias) == 0 {
		gkbase.ErrorCheck(base.ERROR_BAD_REQUEST_MISSING)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_MISSING, base.ERROR_BAD_REQUEST_MISSING
	}

	for _, v := range stream_medias.StreamMedias {
		if v.AccessUUID == "" || v.DomainInner == "" || v.DomainOuter == "" {
			gkbase.ErrorCheck(base.ERROR_BAD_REQUEST_MISSING)
			return http.StatusBadRequest, base.HTTP_BAD_REQUEST_MISSING, base.ERROR_BAD_REQUEST_MISSING
		}
	}

	err = redis_cli.SetStreamMedia(stream_medias.StreamMedias)
	if err != nil {
		gkbase.ErrorCheck(err)
		return http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR_SET_STREAM_MEDIA, err
	}

	return http.StatusCreated, base.HTTP_OK, nil
}

func stream_media_del(w http.ResponseWriter, r *http.Request) (int, string, error) {
	vars := mux.Vars(r)
	index := vars["index"]
	ok := redis_cli.DelStreamMedia(index)
	if ok {
		return http.StatusOK, base.HTTP_OK, nil
	}

	return http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR_DEL_STREAM_MEDIA, nil
}

func stream_media_put(w http.ResponseWriter, r *http.Request) (int, string, error) {
	vars := mux.Vars(r)
	index := vars["index"]
	r.ParseForm()
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var stream_media vcbase.StreamMedia
	err := decoder.Decode(&stream_media)
	if err != nil {
		gkbase.ErrorCheck(err)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_DECODE, err
	}
	if stream_media.AccessUUID == "" || stream_media.DomainInner == "" || stream_media.DomainOuter == "" {
		gkbase.ErrorCheck(base.ERROR_BAD_REQUEST_MISSING)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_MISSING, base.ERROR_BAD_REQUEST_MISSING
	}
	ok := redis_cli.UpdateStreamMedia(index, &stream_media)
	if ok {
		return http.StatusOK, base.HTTP_OK, nil
	}

	return http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR_UPD_STREAM_MEDIA, nil
}

func StreamMedia(w http.ResponseWriter, r *http.Request) {
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
	case http.MethodPost:
		http_status, internal_status, err = stream_media_post(w, r)
	case http.MethodGet:
		http_status, internal_status, data, err = stream_media_get(w, r)
	case http.MethodDelete:
		http_status, internal_status, err = stream_media_del(w, r)
	case http.MethodPut:
		http_status, internal_status, err = stream_media_put(w, r)
	}

	common_reply(w, http_status, internal_status, data, err)
}
