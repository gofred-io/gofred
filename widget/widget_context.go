package widget

import (
	"syscall/js"
)

type WidgetContext struct {
	Current js.Value
	Doc     js.Value
	Root    Widget
}

var (
	context *WidgetContext
)

func NewContext() *WidgetContext {
	doc := js.Global().Get("document")
	root := doc.Call("getElementById", "root")

	return &WidgetContext{
		Current: root,
		Doc:     doc,
		Root:    Widget(root),
	}
}

func Context() *WidgetContext {
	if context == nil {
		context = NewContext()
	}
	return context
}

func (c *WidgetContext) AppendChild(child Widget) {
	c.Current.Call("appendChild", js.Value(child))
}

func (c *WidgetContext) ClientHeight() int {
	return c.Doc.Get("documentElement").Get("clientHeight").Int()
}

func (c *WidgetContext) ClientWidth() int {
	return c.Doc.Get("documentElement").Get("clientWidth").Int()
}

func (c *WidgetContext) CreateElement(tag string) Widget {
	return Widget(c.Doc.Call("createElement", tag))
}

func (c *WidgetContext) CreateElementNS(namespace string, tag string) Widget {
	return Widget(c.Doc.Call("createElementNS", namespace, tag))
}

func (c *WidgetContext) GetElementByID(id string) *Widget {
	element := c.Doc.Call("getElementById", id)
	if element.IsNull() {
		return nil
	}
	w := Widget(element)
	return &w
}

func (c *WidgetContext) Navigate(path string) {
	js.Global().Get("location").Set("href", path)
}

func (c *WidgetContext) OpenLink(href string) {
	js.Global().Get("window").Call("open", href, "_blank")
}
