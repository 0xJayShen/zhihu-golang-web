package main

import (
	_ "net/http/pprof"
	"github.com/asdfsx/zhihu-golang-web/cmd"
)

func main() {
	cmd.Execute()
}
