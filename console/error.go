package console

import "syscall/js"

var (
	console = js.Global().Get("console")
)

func Error(e string) {
	console.Call("error", e)
}

func Log(e string) {
	console.Call("log", e)
}
