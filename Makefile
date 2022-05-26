PROJECT = xtool

SHELL := /bin/bash
BASEDIR = $(shell pwd)
# golang下载代理方式
PROXY = https://goproxy.cn,direct

# build with version infos
versionDir = "xtool/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)
gitAuthor = $(shell git config --global user.name)

ldflags="-w \
-X ${versionDir}.gitAuthor=${gitAuthor}\
-X ${versionDir}.gitTag=${gitTag} \
-X ${versionDir}.buildDate=${buildDate} \
-X ${versionDir}.gitCommit=${gitCommit} \
-X ${versionDir}.gitTreeState=${gitTreeState}"

goModule=GO111MODULE=on CGO_ENABLED=0 GOPROXY=${PROXY}
# (-tags netgo) for alpine
buildCmd=go build -tags netgo -o ${PROJECT} -v -ldflags ${ldflags} .

all: clean gotool
	${goModule} ${buildCmd}

mac: clean gotool
	${goModule} GOOS=darwin GOARCH=amd64 ${buildCmd}

linux: clean gotool
	${goModule} GOOS=linux GOARCH=amd64 ${buildCmd}

windows: clean gotool
	${goModule} GOOS=windows GOARCH=amd64 ${buildCmd}

test: gotool
	${goModule} go test ./...

clean:
	rm -f ${PROJECT}
	find . -name "[._]*.s[a-w][a-z]" | xargs rm -f {}

gotool:
	${goModule} gofmt -w .
	${goModule} go vet . | grep -v vendor | true

help:
	@echo "make - compile the source code"
	@echo "make mac - compile mac platform"
	@echo "make linux - compile linux platform"
	@echo "make windows - compile windows platform"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make test - run go test"

.PHONY: clean gotool help test


