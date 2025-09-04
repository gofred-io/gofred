package widget

import (
	"fmt"
	"net/url"
	"syscall/js"

	"github.com/gofred-io/gofred/utils"
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

func (c *WidgetContext) AppendChild(child BaseWidget) {
	c.Current.Call("appendChild", js.Value(child.Widget))
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

func (c *WidgetContext) CreateElementWithOptions(tag string, options map[string]any) Widget {
	return Widget(c.Doc.Call("createElement", tag, options))
}

func (c *WidgetContext) CreateElementNS(namespace string, tag string) Widget {
	return Widget(c.Doc.Call("createElementNS", namespace, tag))
}

func (c *WidgetContext) CurrentPath() string {
	return js.Global().Get("window").Get("location").Get("pathname").String()
}

func (c *WidgetContext) GetElementByID(id string) *Widget {
	element := c.Doc.Call("getElementById", id)
	if element.IsNull() {
		return nil
	}
	w := Widget(element)
	return &w
}

func (c *WidgetContext) GoBack() {
	js.Global().Get("window").Get("history").Call("back")
}

func (c *WidgetContext) Navigate(path string) {
	js.Global().Get("window").Get("history").Call("pushState", nil, "", path)
}

func (c *WidgetContext) OnNavigate(callback func(path string)) {
	js.Global().Get("window").Call("addEventListener", "popstate", js.FuncOf(func(this js.Value, args []js.Value) any {
		if callback != nil {
			utils.SafeGo(func() {
				arg := args[0]
				destinationUrl, err := url.Parse(arg.Get("target").Get("location").Get("href").String())
				if err != nil {
					fmt.Println("error parsing destination url", err)
					return
				}
				callback(destinationUrl.Path)
			})
		}
		return nil
	}))
}

func (c *WidgetContext) OpenLink(href string) {
	js.Global().Get("window").Call("open", href, "_blank")
}

func (c *WidgetContext) OnResize(callback func()) {
	js.Global().Get("window").Set("onresize", js.FuncOf(func(this js.Value, args []js.Value) any {
		if callback != nil {
			utils.SafeGo(func() {
				callback()
			})
		}
		return nil
	}))
}

func (c *WidgetContext) PushState(state any, title string, url string) {
	js.Global().Get("window").Get("history").Call("pushState", state, title, url)
}
