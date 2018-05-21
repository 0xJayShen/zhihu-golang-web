package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
	"fmt"
)

type Server struct {
	HTTP_PORT     int
	READ_TIMEOUT  time.Duration
	WRITE_TIMEOUT time.Duration
}

type App struct {
	PAGE_SIZE  int
	JWT_SECRET string
}

type RunMode struct {
	RUN_MODE string
}

var (
	Cfg      *ini.File
	RunMode_ *RunMode
	Server_  *Server
	App_     *App
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}
func LoadBase() {
	RunMode_ = new(RunMode)
	err := Cfg.MapTo(RunMode_)
	if err != nil {
		fmt.Println(err)
	}

}

func LoadServer() {
	Server_ = new(Server)
	err := Cfg.Section("Server").MapTo(Server_)
	if err != nil {

	}
}

func LoadApp() {
	App_ = new(App)
	err :=  Cfg.Section("App").MapTo(App_)
	if err != nil {

	}
}
