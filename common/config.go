package common

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
	Database struct {
		Type        string
		Host        string
		Port        int64
		User        string
		Passwd      string
		TablePrefix string `mapstructure:"table_prefix"`
	}
	Redis struct {
		Host     string
		Port     int64
	}
}
