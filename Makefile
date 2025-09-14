ifeq ($(OS),Windows_NT)
    SHELL := cmd.exe
    .SHELLFLAGS := /C
    BUILD = cmd /C "set GOOS=js&& set GOARCH=wasm&& go build -o server/main.wasm main.go"
else
    SHELL := /bin/sh
    BUILD = GOOS=js GOARCH=wasm go build -o server/main.wasm main.go
endif

all: build

build:
	$(BUILD)

serve:
	go run server/server.go

