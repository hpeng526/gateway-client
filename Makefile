export PATH := $(GOPATH)/bin:$(PATH)
export GO15VENDOREXPERIMENT := 1

all: build

build: bilibili weather router

bilibili:
	env GOOS=darwin GOARCH=amd64 go build -o ./build/bilibili/bili-client_darwin_amd64 ./cmd/bilibili/
	env GOOS=linux GOARCH=386 go build -o ./build/bilibili/bili-client_linux_386 ./cmd/bilibili/
	env GOOS=linux GOARCH=amd64 go build -o ./build/bilibili/bili-client_linux_amd64 ./cmd/bilibili/
	env GOOS=windows GOARCH=386 go build -o ./build/bilibili/bili-client_windows_386.exe ./cmd/bilibili/
	env GOOS=linux GOARCH=mips64 go build -o ./build/bilibili/bili-client_linux_mips64 ./cmd/bilibili/
	env GOOS=linux GOARCH=mips64le go build -o ./build/bilibili/bili-client_linux_mips64le ./cmd/bilibili/
	env GOOS=linux GOARCH=mips go build -o ./build/bilibili/bili-client_linux_mips ./cmd/bilibili/
	env GOOS=linux GOARCH=mipsle go build -o ./build/bilibili/bili-client_linux_mipsle ./cmd/bilibili/

weather:
	env GOOS=darwin GOARCH=amd64 go build -o ./build/weather/weather-client_darwin_amd64 ./cmd/weather/
	env GOOS=linux GOARCH=386 go build -o ./build/weather/weather-client_linux_386 ./cmd/weather/
	env GOOS=linux GOARCH=amd64 go build -o ./build/weather/weather-client_linux_amd64 ./cmd/weather/
	env GOOS=windows GOARCH=386 go build -o ./build/weather/weather-client_windows_386.exe ./cmd/weather/
	env GOOS=linux GOARCH=mips64 go build -o ./build/weather/weather-client_linux_mips64 ./cmd/weather/
	env GOOS=linux GOARCH=mips64le go build -o ./build/weather/weather-client_linux_mips64le ./cmd/weather/
	env GOOS=linux GOARCH=mips go build -o ./build/weather/weather-client_linux_mips ./cmd/weather/
	env GOOS=linux GOARCH=mipsle go build -o ./build/weather/weather-client_linux_mipsle ./cmd/weather/

router:
	env GOOS=linux GOARCH=mipsle go build -o ./build/router/router_mipsle ./cmd/router/

clean:
	rm -rf ./build/