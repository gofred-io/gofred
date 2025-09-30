package application

import (
	"syscall/js"

	"github.com/gofred-io/gofred/utils"
)

var (
	Nil = BaseWidget{Widget: Widget(js.Value{})}
)

type Widget js.Value

func (w Widget) AddClass(class string) {
	w.Get("classList").Call("add", class)
}

func (w Widget) AppendChild(child Widget) {
	if child.IsNull() || child.Equal(Nil) {
		return
	}

	w.Call("appendChild", js.Value(child))
}

func (w Widget) Call(m string, args ...any) Widget {
	return Widget(js.Value(w).Call(m, args...))
}

func (w Widget) ChildAt(index int) Widget {
	return Widget(js.Value(w).Get("children").Index(index))
}

func (w Widget) Equal(other BaseWidget) bool {
	return js.Value(w).Equal(js.Value(other.Widget))
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

func (w Widget) Remove() {
	w.Call("remove")
}

func (w Widget) RemoveClass(class string) {
	w.Get("classList").Call("remove", class)
}

func (w Widget) ReplaceWith(widget Widget) {
	w.Call("replaceWith", js.Value(widget))
}

func (w Widget) Set(p string, v any) {
	js.Value(w).Set(p, v)
}

func (w Widget) SetAttribute(p string, v any) {
	js.Value(w).Call("setAttribute", p, v)
}

func (w Widget) SetAttributeNS(namespace string, p string, v any) {
	js.Value(w).Call("setAttributeNS", namespace, p, v)
}

func (w Widget) SetClass(class string) {
	w.Set("className", class)
}

func (w Widget) SetDraggable(draggable bool) {
	w.Set("draggable", draggable)
}

func (w Widget) SetID(id string) {
	w.Set("id", id)
}

func (w Widget) SetOnClick(onClick func(this BaseWidget, e Event)) {
	w.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) any {
		if onClick != nil {
			utils.SafeGo(func() {
				e := args[0]
				onClick(BaseWidget{Widget: Widget(this)}, Event(e))
			})
		}
		return nil
	}))
}

func (w Widget) SetOnDragEnd(onDragEnd func(this BaseWidget, e Event)) {
	w.Set("ondragend", js.FuncOf(func(this js.Value, args []js.Value) any {
		if onDragEnd != nil {
			utils.SafeGo(func() {
				e := args[0]
				onDragEnd(BaseWidget{Widget: Widget(this)}, Event(e))
			})
		}
		return nil
	}))
}

func (w Widget) SetOnDrag(onDrag func(this BaseWidget, e Event)) {
	w.Set("ondrag", js.FuncOf(func(this js.Value, args []js.Value) any {
		if onDrag != nil {
			utils.SafeGo(func() {
				e := args[0]
				onDrag(BaseWidget{Widget: Widget(this)}, Event(e))
			})
		}
		return nil
	}))
}

func (w Widget) SetOnDragStart(onDragStart func(this BaseWidget, e Event)) {
	w.Set("ondragstart", js.FuncOf(func(this js.Value, args []js.Value) any {
		if onDragStart != nil {
			utils.SafeGo(func() {
				e := args[0]
				onDragStart(BaseWidget{Widget: Widget(this)}, Event(e))
			})
		}
		return nil
	}))
}

func (w Widget) SetStyle(style string) {
	w.Set("style", style)
}

func (w Widget) ToggleClass(class string) {
	w.Get("classList").Call("toggle", class)
}

func (w Widget) UpdateStyleProperty(property string, value string) {
	w.Get("style").Set(property, value)
}

func (w Widget) SetText(text string) {
	w.Set("textContent", text)
}
