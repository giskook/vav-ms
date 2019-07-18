package v1

import (
	myhttp "github.com/giskook/go/http"
	"github.com/giskook/vav-ms/base"
	"net/http"
)

func common_reply(w http.ResponseWriter, http_status int, code string, data interface{}, err error) {
	w.WriteHeader(http_status)
	err_msg := ""
	if err != nil {
		err_msg = err.Error()
	} else {
		err_msg = base.ErrorMap[code]
	}
	myhttp.EncodeResponse(w, code, data, err_msg)
}
