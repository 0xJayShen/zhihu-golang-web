package cmd

import (
	"testing"
	"github.com/prashantv/gostub"
	"github.com/smartystreets/goconvey/convey"
)

func TestReadconfig(t *testing.T){
	stubs := gostub.New()
	stubs.Stub(&config, "../server.yml")
	convey.Convey("readconfig", t, func(){
		convey.So(readConfig(), convey.ShouldBeNil)
	})
	convey.Convey("test level", t, func(){
		convey.So(config.Log.Level, convey.ShouldEqual, "info")
	})
	convey.Convey("test level", t, func(){
		convey.So(config.Log.Path, convey.ShouldEqual, "/tmp/server")
	})
}

