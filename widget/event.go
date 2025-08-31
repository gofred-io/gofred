package widget

import "syscall/js"

type Event js.Value

func (e Event) PreventDefault() {
	js.Value(e).Call("preventDefault")
}

func (e Event) StopPropagation() {
	js.Value(e).Call("stopPropagation")
}
