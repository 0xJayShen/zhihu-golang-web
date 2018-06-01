package cmd

import (
	"testing"
	"github.com/prashantv/gostub"
	"github.com/smartystreets/goconvey/convey"
)

func TestReadconfig(t *testing.T){
	stubs := gostub.New()
	stubs.Stub(&cfgFile, "../server.yml")
	defer stubs.Reset()

	convey.Convey("readconfig", t, func(){
		convey.So(readConfig(), convey.ShouldBeNil)
	})
	convey.Convey("test level", t, func(){
		convey.So(config.Log.Level, convey.ShouldEqual, "info")
	})
	convey.Convey("test log Path", t, func(){
		convey.So(config.Log.Path, convey.ShouldEqual, "/tmp/server")
	})

	convey.Convey("test pprof", t, func(){
		convey.So(config.Pprof.Listen, convey.ShouldEqual, "0.0.0.0:6060")
	})

	convey.Convey("test etcd", t, func(){
		convey.So(config.Etcd.Path, convey.ShouldEqual, "/server")
	})

	convey.Convey("test database", t, func(){
		convey.So(config.Database.TablePrefix, convey.ShouldEqual, "shop_")
	})

	convey.Convey("test redis", t, func(){
		convey.So(config.Redis.Port, convey.ShouldEqual, 6389)
	})
}

