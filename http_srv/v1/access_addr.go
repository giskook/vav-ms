package v1

import (
	"encoding/json"
	gkbase "github.com/giskook/go/base"
	gkhttp "github.com/giskook/go/http"
	"github.com/giskook/vav-ms/base"
	"github.com/giskook/vav-ms/redis_cli"
	"net/http"
)

type access_addr struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func access_addr_get(w http.ResponseWriter, r *http.Request) (int, string, interface{}, error) {
	ip, port, err := redis_cli.GetAccessAddr()
	if err != nil {
		return http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR_SET_ACCESS_ADDR, nil, err
	}

	return http.StatusOK, base.HTTP_OK,
		&access_addr{IP: ip,
			Port: port,
		}, nil
}

func access_addr_post(w http.ResponseWriter, r *http.Request) (int, string, error) {
	r.ParseForm()
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var addr access_addr
	err := decoder.Decode(&addr)
	if err != nil {
		gkbase.ErrorCheck(err)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_DECODE, err
	}
	if addr.IP == "" ||
		addr.Port == "" {
		gkbase.ErrorCheck(base.ERROR_BAD_REQUEST_MISSING)
		return http.StatusBadRequest, base.HTTP_BAD_REQUEST_MISSING, base.ERROR_BAD_REQUEST_MISSING
	}

	err = redis_cli.SetAccessAddr(addr.IP, addr.Port)
	if err != nil {
		gkbase.ErrorCheck(err)
		return http.StatusInternalServerError, base.HTTP_INTERNAL_SERVER_ERROR_SET_ACCESS_ADDR, err
	}

	return http.StatusCreated, base.HTTP_OK, nil
}

func AccessAddr(w http.ResponseWriter, r *http.Request) {
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
		http_status, internal_status, err = access_addr_post(w, r)
	case http.MethodGet:
		http_status, internal_status, data, err = access_addr_get(w, r)
	}

	common_reply(w, http_status, internal_status, data, err)
}
