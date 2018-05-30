package cmd

import (
	"os"
	"os/signal"
	"time"
	"context"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/cobra"
	"github.com/asdfsx/zhihu-golang-web/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "web server's main process",
	Long:  "read the config file, read logs according to the configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := readConfig()
		if err != nil{
			return err
		}
		err = serve(cmd, args)
		if err == nil{
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.server.yaml)")
}

func logrotate(){
	//判断是否应该rotate日志
	now := time.Now()
	if now.Day() > rotateTime.Day() {
		if err := logger.Rotate(); err != nil {
			jww.ERROR.Printf("Failed to rotate the logger %v\n", err)
		}
		rotateTime = now
	}
}

func serve(cmd *cobra.Command, args []string) error {
	jww.INFO.Printf("server name: %v\n", config.Server.Name)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer close(signalChan)

	ticker := time.Tick(5 * time.Minute)

	svr := server.NewServer(&config)
	defer svr.Close()

	err := svr.Serv(ctx)
	if err != nil {
		jww.ERROR.Println(err)
		return err
	}

	for {
		select {
			case <- signalChan:
				cancel()
				return nil
			case <-ticker:
				logrotate()
		}
	}
}
