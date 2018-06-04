package main

import (
	"fmt"
	"net/http"

	"github.com/qq976739120/zhihu-golang-web/pkg/setting"
	"github.com/qq976739120/zhihu-golang-web/routers"
	"time"
	"log"
	"os"
	"os/signal"
	"context"
	"go_partice/tail_kfka/settings"
	"github.com/qq976739120/zhihu-golang-web/queue"
	"github.com/qq976739120/zhihu-golang-web/pkg/util"
	"github.com/qq976739120/zhihu-golang-web/elastic"
	"github.com/qq976739120/zhihu-golang-web/pkg/logging"
)

func main() {
	logging.Info("do it")
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Server_.HTTP_PORT),
		Handler:        router,
		ReadTimeout:    setting.Server_.READ_TIMEOUT * time.Second,
		WriteTimeout:   setting.Server_.WRITE_TIMEOUT * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := util.InitTail(settings.Collect__.Collectlist, settings.Collect_.ChanSize)
	err = queue.InitKafka(settings.Kafka_.KafkaAddress)
	err = elastic.InitES(settings.Elastic_.ESaddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	go queue.KafkaToES()
	go queue.TailToKafka()

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
