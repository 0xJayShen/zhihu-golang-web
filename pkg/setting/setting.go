package setting

import (
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
	PAGE_SIZE       int
	JWT_SECRET      string
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
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
	PassWord         string
}

type Kafka struct {
	KafkaAddress string
	Topic        string
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
	RunMode_  = &RunMode{}
	Server_   = &Server{}
	App_      = &App{}
	DataBase_ = &DataBase{}
	Redis_    = &Redis{}
	Kafka_    = &Kafka{}

	Collect_ = &Collect{}
	Elastic_ = Elastic{}
)

func InitSettings() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()

	mapTo("App", App_)
	mapTo("Server", Server_)
	mapTo("DataBase", DataBase_)
	mapTo("Redis", Redis_)
	mapTo("Collect", Collect_)
	mapTo("Kafka", Kafka_)
	mapTo("Elastic", Elastic_)

	Server_.READ_TIMEOUT = Server_.READ_TIMEOUT * time.Second
	Server_.WRITE_TIMEOUT = Server_.WRITE_TIMEOUT * time.Second
	//伪造的 collect 切片
	Collect__.Collectlist = append(Collect__.Collectlist, *Collect_)
}

func LoadBase() {
	RunMode_ = new(RunMode)
	err := Cfg.MapTo(RunMode_)
	if err != nil {
		fmt.Println("初始化 base失败")
	}
}

func mapTo(section string, v interface{}) {
	err := Cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Println(err, "初始化配置文件")
	}
}
