package main

import (
	"fmt"
	"net/http"


	"zhihu-golang-web/pkg/setting"
	"zhihu-golang-web/routers"

	"reflect"
	"time"
)

func main() {
	router := routers.InitRouter()
	fmt.Println(setting.Server_.READ_TIMEOUT,reflect.TypeOf(setting.Server_.READ_TIMEOUT))
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Server_.HTTP_PORT),
		Handler:        router,
		ReadTimeout:    setting.Server_.READ_TIMEOUT * time.Second,
		WriteTimeout:   setting.Server_.WRITE_TIMEOUT * time.Second,
		MaxHeaderBytes: 1 << 20,
	}


	err :=s.ListenAndServe()
	fmt.Println(err)
}
