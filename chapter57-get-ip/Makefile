arch=amd64
archARM=arm64

build:
	GOOS=linux GOARCH=$(arch) go build -o bin/getip ./main.go

build-arm:
	GOOS=linux GOARCH=$(archARM) go build -o bin/getip-arm ./main.go
