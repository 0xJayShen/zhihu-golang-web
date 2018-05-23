package main

import (
	"fmt"
	"net/http"

	"zhihu-golang-web/pkg/setting"
	"zhihu-golang-web/routers"

	"reflect"
	"time"
	"log"
	"os"
	"os/signal"
	"context"
)

func main() {
	router := routers.InitRouter()
	fmt.Println(setting.Server_.READ_TIMEOUT, reflect.TypeOf(setting.Server_.READ_TIMEOUT))
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Server_.HTTP_PORT),
		Handler:        router,
		ReadTimeout:    setting.Server_.READ_TIMEOUT * time.Second,
		WriteTimeout:   setting.Server_.WRITE_TIMEOUT * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
