package server

//all configs inside config file
type Config struct {
	Server struct{
		Name         string
		Port         uint64
		ReadTimeout  uint64
		WriteTimeout uint64
	}
	Log struct {
		Level string
		Path  string
	}
	Pprof struct {
		Enable bool
		Listen string
	}
	Etcd struct {
		Enable   bool
		Addr     string
		Path     string
		LeaseTTL int64
	}
	Mysql struct {
		Host     string
		Port     int64
		User     string
		Passwd   string
	}
	Redis struct {
		Host     string
		Port     int64
	}
}
