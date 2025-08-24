all: build

build:
	GOARCH=wasm GOOS=js go build -o server/main.wasm main.go

serve:
	go run server/server.go