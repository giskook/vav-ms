package http_srv

import (
	"context"
	"github.com/giskook/vav-ms/conf"
	"github.com/giskook/vav-ms/http_srv/v1"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HttpSrv struct {
	conf   *conf.Conf
	router *mux.Router
	srv    *http.Server
}

func NewHttpSrv(conf *conf.Conf) *HttpSrv {
	return &HttpSrv{
		conf:   conf,
		router: mux.NewRouter(),
	}
}

func (h *HttpSrv) Start() {
	h.router.StrictSlash(true)
	s := h.router.PathPrefix("/vavms/api").Subrouter()
	h.init_api_v1(s)

	h.srv = &http.Server{
		Addr: h.conf.Http.Addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: h.conf.Http.WriteTimeOut,
		ReadTimeout:  h.conf.Http.ReadTimeOut,
		IdleTimeout:  h.conf.Http.IdleTimeOut,
		Handler:      h.router, // Pass our instance of gorilla/mux in.
	}

	log.Printf("<INF> http listening %s\n", h.conf.Http.Addr)
	go func() {
		if err := h.srv.ListenAndServe(); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}

func (h *HttpSrv) init_api_v1(r *mux.Router) {
	s := r.PathPrefix("/v1").Subrouter()
	s.HandleFunc("/stream_media/{index}", v1.StreamMedia).Methods("DELETE", "PUT")
	s.HandleFunc("/stream_media", v1.StreamMedia).Methods("POST", "GET")
	s.HandleFunc("/access_addr", v1.AccessAddr).Methods("POST", "GET")
	//s.HandleFunc("/user/{user_id}/sms", v1.Handler_api_sms).Methods("POST")
}

func (h *HttpSrv) ShutDown() {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), h.conf.Http.ShutWaitTimeOut)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	h.srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
}
