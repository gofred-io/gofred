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

func (w Widget) ChildAt(index int) Widget {
	return Widget(js.Value(w).Get("children").Index(index))
}

func (w Widget) Get(p string) js.Value {
	return js.Value(w).Get(p)
}

func (w Widget) ID() string {
	return w.Get("id").String()
}

func (w Widget) IsNull() bool {
	return js.Value(w).IsNull()
}

func (w Widget) Parent() Widget {
	return Widget(js.Value(w).Get("parentElement"))
}

func (w Widget) Set(p string, v any) {
	js.Value(w).Set(p, v)
}

func (w Widget) SetID(id string) {
	w.Set("id", id)
}

func (w Widget) SetOnClick(onClick func(this Widget)) {
	_onClick := func(this Widget) any {
		if onClick != nil {
			onClick(this)
		}

		return nil
	}

	w.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) any {
		return _onClick(Widget(this))
	}))
}

func (w Widget) SetStyle(style string) {
	w.Set("style", style)
}

func (w Widget) SetText(text string) {
	w.Set("textContent", text)
}
