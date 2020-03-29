.PHONY: list fmt 
all: list fmt 
BINARY="utils"
VERSION=0.0.1
BUILD=`date +%F`

PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default:
	@echo "build the ${BINARY}"
	@#这个--tags=jsoniter是个什么鬼啊
	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ${BINARY}  -tags=jsoniter
	@echo "build done."

list:
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

fmt:
	@echo "fmt the project"
	@gofmt -s -w ${GOFILES}

