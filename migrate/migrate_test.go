package migrate

import (
	"testing"
	"github.com/prashantv/gostub"
    "github.com/smartystreets/goconvey/convey"
)

var dbType, user, password, host, dbName, tablePrefix string
var port int64

func TestReadconfig(t *testing.T){
	stubs := gostub.New()
	stubs.Stub(&dbType, "mysql")
	stubs.Stub(&user, "root")
	stubs.Stub(&password, "root")
	stubs.Stub(&host, "127.0.0.1")
	stubs.Stub(&port, "3306")
	stubs.Stub(&dbName, "shop")
	stubs.Stub(&tablePrefix, "shop_")
	defer stubs.Reset()

	convey.Convey("create connection", t, func(){
		convey.So(ConnectDB(dbType, user, password, host, port, dbName, tablePrefix), convey.ShouldBeNil)
	})

	convey.Convey("migrate", t, func(){
		convey.So(Migrate(), convey.ShouldBeNil)
	})

	convey.Convey("close", t, func(){
		convey.So(Close(), convey.ShouldBeNil)
	})
}

