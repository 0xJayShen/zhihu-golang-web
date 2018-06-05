package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/qq976739120/zhihu-golang-web/pkg/file"
	"github.com/qq976739120/zhihu-golang-web/pkg/setting"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.App_.RuntimeRootPath, setting.App_.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.App_.LogSaveName,
		time.Now().Format(setting.App_.TimeFormat),
		setting.App_.LogFileExt,
	)
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
