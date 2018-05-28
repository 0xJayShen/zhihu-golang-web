package server

import (
	"net/http"
	"context"
	jww "github.com/spf13/jwalterweatherman"
	"fmt"
	"time"
	"github.com/asdfsx/zhihu-golang-web/routers"
)

type Server struct {
	config *Config
    svr    *http.Server
}

func NewServer(config *Config) *Server {
	router := routers.InitRouter()
	return &Server{config,
	&http.Server{
		Addr:           fmt.Sprintf(":%d", config.Server.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(config.Server.ReadTimeout),
		WriteTimeout:   time.Duration(config.Server.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}}
}

func(server *Server) Close() error {
	return server.svr.Close()
}

func(server *Server) Serv (ctx context.Context) error {
	jww.INFO.Println("start serv...")
	return server.svr.ListenAndServe()
}