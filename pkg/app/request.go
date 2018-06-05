package app

import "github.com/qq976739120/zhihu-golang-web/pkg/logging"

func MarkErrors(errors []error){
	for _,err := range errors{
		logging.Info(err)
	}
}