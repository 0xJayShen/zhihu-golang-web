package cmd

import (
	"fmt"
	"time"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/asdfsx/lumberjack"
	"github.com/spf13/viper"
	"github.com/asdfsx/zhihu-golang-web/common"
)

var (
	cfgFile string
	config common.Config
	logger  *lumberjack.Logger
	rotateTime time.Time
)

// readConfig reads in config file and ENV variables if set.
func readConfig() error {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("server") // name of config file (without extension)
		viper.AddConfigPath("$HOME")         // adding home directory as first search path
	}

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		jww.ERROR.Println(err)
		return err
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())

	err = viper.Unmarshal(&config)
	if err != nil {
		jww.ERROR.Println(err)
		return err
	}

	switch config.Log.Level {
	case "trace":
		jww.SetLogThreshold(jww.LevelTrace)
	case "info":
		jww.SetLogThreshold(jww.LevelInfo)
	case "debug":
		jww.SetLogThreshold(jww.LevelDebug)
	default:
		jww.SetLogThreshold(jww.LevelWarn)
	}

	if config.Log.Path != "" {
		logger = &lumberjack.Logger{
			Filename:         config.Log.Path + "/kafkaproducer.log",
			MaxSize:          500, // megabytes
			MaxBackups:       3,
			MaxAge:           3, //days
			LocalTime:        true,
			BackupTimeFormat: "20060102T150405",
		}
		jww.SetLogOutput(logger)
	}
	return nil
}
