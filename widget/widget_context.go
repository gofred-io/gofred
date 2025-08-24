package widget

import (
	"syscall/js"
)

type WidgetContext struct {
	Current js.Value
	Doc     js.Value
	Root    js.Value
}

func Context() *WidgetContext {
	doc := js.Global().Get("document")
	root := doc.Call("getElementById", "root")

	return &WidgetContext{
		Current: root,
		Doc:     doc,
		Root:    root,
	}
}

func (c *WidgetContext) AppendChild(child Widget) {
	c.Current.Call("appendChild", js.Value(child))
}

func (c *WidgetContext) CreateElement(tag string) Widget {
	return Widget(c.Doc.Call("createElement", tag))
}
