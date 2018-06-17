PACKAGE = github.com/asdfsx/zhihu-golang-web
COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE = `date +%FT%T%z`
LDFLAGS = -ldflags "-X ${PACKAGE}/server.CommitHash=${COMMIT_HASH} -X ${PACKAGE}/server.BuildDate=${BUILD_DATE}"
NOGI_LDFLAGS = -ldflags "-X ${PACKAGE}/server.BuildDate=${BUILD_DATE}"

.PHONY: server, vendor

vendor: ## Install dep and sync server's vendored dependencies
	go get -u github.com/golang/dep/cmd/dep
	export http_proxy=http://127.0.0.1:1087;export https_proxy=http://127.0.0.1:1087; dep ensure

server1: 
	go build ${LDFLAGS} ${PACKAGE}

server: vendor ## Build server binary
	go build ${LDFLAGS} ${PACKAGE}

server4linux: vendor ## cross compile 4 linux
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} ${PACKAGE}

server-race: vendor ## Build server binary with race detector enabled
	go build -race ${LDFLAGS} ${PACKAGE}

install: vendor ## Install server binary
	go install ${LDFLAGS} ${PACKAGE}

clean: ## delete the build target
	rm zhihu-golang-web

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
