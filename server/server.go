package server

import (
	"net/http"
	"context"
	jww "github.com/spf13/jwalterweatherman"
	"fmt"
	"time"
	"github.com/asdfsx/zhihu-golang-web/routers"
	"github.com/asdfsx/zhihu-golang-web/common"
	"github.com/jinzhu/gorm"
)

type Server struct {
	config                        *common.Config
    svr                           *http.Server
    db                            *gorm.DB
	stopChan                      chan bool
}

func NewServer(config *common.Config) (*Server, error) {
	router := routers.InitRouter()
	db, err := connect(config.Database.Type, config.Database.User, config.Database.Passwd,
		config.Database.Host, config.Database.Port, config.Database.DBName, config.Database.TablePrefix)
	if err != nil{
		return nil, err
	}
	server := new(Server)
	server.config = config
	server.svr = &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Server.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(config.Server.ReadTimeout),
		WriteTimeout:   time.Duration(config.Server.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}
	server.db = db

	return server, nil
}

func(server *Server) Close(){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go func(ctx context.Context) {
		<-ctx.Done()
		if ctx.Err() == context.Canceled {
			return
		} else if ctx.Err() == context.DeadlineExceeded {
			panic("Timeout while stopping server, killing instance ✝")
		}
	}(ctx)
	server.svr.Close()
	server.db.Close()
	cancel()
}

func(server *Server) Start () error {
	jww.INFO.Println("start serv...")
	return server.svr.ListenAndServe()
}

func (s *Server) StartWithContext(ctx context.Context) {
	go func() {
		defer s.Close()
		<-ctx.Done()
		jww.INFO.Println("I have to go...")
		reqAcceptGraceTimeOut := time.Duration(1000)
		if reqAcceptGraceTimeOut > 0 {
			jww.INFO.Printf("Waiting %s for incoming requests to cease", reqAcceptGraceTimeOut)
			time.Sleep(reqAcceptGraceTimeOut)
		}
		jww.INFO.Println("Stopping server gracefully")
		s.Close()
	}()
	go s.Start()
}

// Wait blocks until server is shutted down.
func (s *Server) Wait() {
	<-s.stopChan
}

// Stop stops the server
func (s *Server) Stop() {
	defer jww.INFO.Println("Server stopped")
	s.stopChan <- true
}

func connect(dbType, user, password, host string, port int64, dbName, tablePrefix string) (db *gorm.DB, err error) {
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil{
		jww.ERROR.Println(err)
		return nil, err
	}
	return db, nil
}
