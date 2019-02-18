package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Result response result to client
type Result struct {
	err  int16
	data []byte
}

//CallBack request call back
type CallBack func(*Result)

// IHandler 调度器接口, 逻辑对接.
type IHandler interface {
	OnReq(data []byte, cb CallBack)
}

// Server http object
type Server struct {
	url    string
	hander IHandler
}

//NewServer http server
func NewServer(url string, hand *IHandler) *Server {
	return &Server{
		url:    url,
		hander: *hand,
	}
}

//Start http start
func (s *Server) Start() {
	srv := http.Server{
		Addr:    s.url,
		Handler: s,
	}

	log.Printf("http server listen at:%s", s.url)
	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("http server ListenAndServe %s failed. %v \r\n", s.url, err)
	}
}

//ServerHTTP implement handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("http server ServeHTTP %s handler failed. %v", s.url, err)
		return
	}

	s.hander.OnReq(data, func(rs *Result) {
		if rs != nil {
			if rs.err == 0 {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
			w.Write(rs.data)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
}
