arch=amd64
archARM=arm64

build:
	GOOS=linux GOARCH=$(arch) go build -o bin/gotest ./main.go

build-arm:
	GOOS=linux GOARCH=$(archARM) go build -o bin/gotest-arm ./main.go
