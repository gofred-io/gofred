package widget

import (
	"syscall/js"
)

type Widget js.Value

func (w Widget) AppendChild(child Widget) {
	w.Call("appendChild", js.Value(child))
}

func (w Widget) Call(m string, args ...any) Widget {
	return Widget(js.Value(w).Call(m, args...))
}

func (w Widget) SetID(id string) {
	w.Set("id", id)
}

func (w Widget) SetStyle(style string) {
	w.Set("style", style)
}

func (w Widget) Set(p string, v any) {
	js.Value(w).Set(p, v)
}
