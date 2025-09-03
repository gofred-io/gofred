package utils

import "syscall/js"

func CopyToClipboard(text string) {
	js.Global().Get("navigator").Get("clipboard").Call("writeText", text)
}
