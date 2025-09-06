package widget

import (
	"syscall/js"

	"github.com/gofred-io/gofred/utils"
)

var (
	Nil = JSWidget(js.Value{})
)

type JSWidget js.Value

func (w *JSWidget) AddClass(class string) {
	w.Get("classList").Call("add", class)
}

func (w *JSWidget) AppendChild(child Widget) {
	if child == nil || child.GetBaseWidget().JSWidget.IsNull() {
		return
	}

	w.Call("appendChild", js.Value(*child.GetBaseWidget().JSWidget))
}

func (w *JSWidget) Call(m string, args ...any) JSWidget {
	return JSWidget(js.Value(*w).Call(m, args...))
}

func (w *JSWidget) ChildAt(index int) JSWidget {
	return JSWidget(js.Value(*w).Get("children").Index(index))
}

func (w *JSWidget) Equal(other BaseWidget) bool {
	return js.Value(*w).Equal(js.Value(*other.JSWidget))
}

func (w *JSWidget) Get(p string) js.Value {
	return js.Value(*w).Get(p)
}

func (w *JSWidget) ID() string {
	return (*w).Get("id").String()
}

func (w *JSWidget) IsNull() bool {
	return js.Value(*w).IsNull()
}

func (w *JSWidget) Parent() JSWidget {
	return JSWidget(js.Value(*w).Get("parentElement"))
}

func (w *JSWidget) Remove() {
	w.Call("remove")
}

func (w *JSWidget) RemoveClass(class string) {
	w.Get("classList").Call("remove", class)
}

func (w *JSWidget) RemoveStyleProperty(property string) {
	w.Get("style").Set(property, "")
}

func (w *JSWidget) ReplaceWith(widget Widget) {
	if widget == nil || widget.GetBaseWidget().JSWidget.IsNull() {
		w.Call("replaceWith", js.Value(Nil))
		return
	}

	w.Call("replaceWith", js.Value(*widget.GetBaseWidget().JSWidget))
}

func (w *JSWidget) Set(p string, v any) {
	js.Value(*w).Set(p, v)
}

func (w *JSWidget) SetAttribute(p string, v any) {
	js.Value(*w).Call("setAttribute", p, v)
}

func (w *JSWidget) SetAttributeNS(namespace string, p string, v any) {
	js.Value(*w).Call("setAttributeNS", namespace, p, v)
}

func (w *JSWidget) SetClass(class string) {
	w.Set("className", class)
}

func (w *JSWidget) SetID(id string) {
	w.Set("id", id)
}

func (w *JSWidget) SetOnClick(onClick func(this Widget, e Event)) {
	w.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) any {
		if onClick != nil {
			utils.SafeGo(func() {
				e := args[0]
				jsWidget := JSWidget(this)
				w := &BaseWidget{JSWidget: &jsWidget}
				onClick(w, Event(e))
			})
		}
		return nil
	}))
}

func (w *JSWidget) SetStyle(style string) {
	w.Set("style", style)
}

func (w *JSWidget) ToggleClass(class string) {
	w.Get("classList").Call("toggle", class)
}

func (w *JSWidget) UpdateStyleProperty(property string, value string) {
	w.Get("style").Set(property, value)
}

func (w *JSWidget) SetText(text string) {
	w.Set("textContent", text)
}
