export PATH := $(GOPATH)/bin:$(PATH)
export GO15VENDOREXPERIMENT := 1

all: build

build: client

client:
	env GOOS=darwin GOARCH=amd64 go build -o ./build/bilibili/bili-client_darwin_amd64 ./cmd/bilibili/
	env GOOS=linux GOARCH=386 go build -o ./build/bilibili/bili-client_linux_386 ./cmd/bilibili/
	env GOOS=linux GOARCH=amd64 go build -o ./build/bilibili/bili-client_linux_amd64 ./cmd/bilibili/
	env GOOS=windows GOARCH=386 go build -o ./build/bilibili/bili-client_windows_386.exe ./cmd/bilibili/
	env GOOS=linux GOARCH=mips64 go build -o ./build/bilibili/bili-client_linux_mips64 ./cmd/bilibili/
	env GOOS=linux GOARCH=mips64le go build -o ./build/bilibili/bili-client_linux_mips64le ./cmd/bilibili/
	env GOOS=linux GOARCH=mips go build -o ./build/bilibili/bili-client_linux_mips ./cmd/bilibili/
	env GOOS=linux GOARCH=mipsle go build -o ./build/bilibili/bili-client_linux_mipsle ./cmd/bilibili/

clean:
	rm -rf ./build/