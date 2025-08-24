all: build

build:
	tinygo build -o server/main.wasm -target=wasm main.go

serve:
	goexec "http.ListenAndServe(\":3000\", http.FileServer(http.Dir(\"./server\")))"