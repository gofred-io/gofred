all: build

build:
	GOARCH=wasm GOOS=js go build -o server/main.wasm main.go

serve:
	goexec "http.ListenAndServe(\":3000\", http.FileServer(http.Dir(\"./server\")))"