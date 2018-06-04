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
type DataBase struct {
	TYPE         string
	USER         string
	PASSWORD     string
	HOST         string
	NAME         string
	TABLE_PREFIX string
}
type Redis struct {
	RedisAddress     string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
}

type Kafka struct {
	KafkaAddress string
	Topic        string
}
type Logs struct {
	LogPath string
}
type Collect struct {
	LogPath  string
	Topic    string
	ChanSize int
}
type Elastic struct {
	ESaddress string
}

type CollectList struct {
	Collectlist [] Collect
}

var Collect__ CollectList

var (
	Cfg       *ini.File
	RunMode_  *RunMode
	Server_   *Server
	App_      *App
	DataBase_ *DataBase
	Redis_    *Redis
	Kafka_    *Kafka
	Logs_     *Logs
	Collect_  *Collect
	Elastic_  *Elastic
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
	LoadDataBase()
	LoadRedis()
	LoadLogs()
	LoadCollect()
	LoadKafka()
	LoadElastic()
	//伪造的 collect 切片
	Collect__.Collectlist = append(Collect__.Collectlist, *Collect_)
}
func LoadBase() {
	RunMode_ = new(RunMode)
	err := Cfg.MapTo(RunMode_)
	if err != nil {

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
	err := Cfg.Section("App").MapTo(App_)
	if err != nil {

	}
}
func LoadDataBase() {
	DataBase_ = new(DataBase)
	err := Cfg.Section("DataBase").MapTo(DataBase_)
	if err != nil {

	}
}

func LoadRedis() {
	Redis_ = new(Redis)
	err := Cfg.Section("Redis").MapTo(Redis_)
	if err != nil {

	}
}
func LoadKafka() {
	Kafka_ = new(Kafka)
	err := Cfg.Section("Kafka").MapTo(Kafka_)
	if err != nil {

	}
}
func LoadLogs() {
	Logs_ = new(Logs)
	err := Cfg.Section("Logs").MapTo(Logs_)
	if err != nil {
		fmt.Println(err)
	}
}

func LoadCollect() {
	Collect_ = new(Collect)
	err := Cfg.Section("Collect").MapTo(Collect_)
	if err != nil {

	}
}

func LoadElastic() {
	Elastic_ = new(Elastic)
	err := Cfg.Section("Elastic").MapTo(Elastic_)
	if err != nil {

	}
}
