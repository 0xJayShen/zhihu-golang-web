package server

import (
	"fmt"
	"html/template"

	"github.com/asdfsx/zhihu-golang-web/helpers"
)

var (
	// CommitHash contains the current Git revision. Use make to build to make
	// sure this gets set.
	CommitHash string

	// BuildDate contains the date of the current build.
	BuildDate string
)

var kafkaproducerInfo *KafkaproducerInfo

// HugoInfo contains information about the current Hugo environment
type KafkaproducerInfo struct {
	Version    string
	Generator  template.HTML
	CommitHash string
	BuildDate  string
}

func init() {
	kafkaproducerInfo = &KafkaproducerInfo{
		Version:    helpers.ServerVersion(),
		CommitHash: CommitHash,
		BuildDate:  BuildDate,
		Generator:  template.HTML(fmt.Sprintf(`<meta name="generator" content="Hugo %s" />`, helpers.ServerVersion())),
	}
}
