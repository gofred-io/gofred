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
	Root    JSWidget
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
		Root:    JSWidget(root),
	}
}

func Context() *WidgetContext {
	if context == nil {
		context = NewContext()
	}
	return context
}

func (c *WidgetContext) AppendChild(child Widget) {
	c.Root.Call("appendChild", js.Value(*child.GetBaseWidget().JSWidget))
}

func (c *WidgetContext) ClientHeight() int {
	return c.Doc.Get("documentElement").Get("clientHeight").Int()
}

func (c *WidgetContext) ClientWidth() int {
	return c.Doc.Get("documentElement").Get("clientWidth").Int()
}

func (c *WidgetContext) CreateElement(tag string) *JSWidget {
	w := JSWidget(c.Doc.Call("createElement", tag))
	return &w
}

func (c *WidgetContext) CreateElementWithOptions(tag string, options map[string]any) *JSWidget {
	w := JSWidget(c.Doc.Call("createElement", tag, options))
	return &w
}

func (c *WidgetContext) CreateElementNS(namespace string, tag string) *JSWidget {
	w := JSWidget(c.Doc.Call("createElementNS", namespace, tag))
	return &w
}

func (c *WidgetContext) CurrentPath() string {
	return js.Global().Get("window").Get("location").Get("pathname").String()
}

func (c *WidgetContext) GetElementByID(id string) *JSWidget {
	element := c.Doc.Call("getElementById", id)
	if element.IsNull() {
		return nil
	}
	w := JSWidget(element)
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
