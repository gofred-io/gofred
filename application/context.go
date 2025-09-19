package application

import (
	"fmt"
	"net/url"
	"syscall/js"

	"github.com/gofred-io/gofred/utils"
)

type AppContext struct {
	Current  js.Value
	Doc      js.Value
	Root     Widget
	Scaffold any // scaffold.Scaffold
}

var (
	context *AppContext = NewContext()
)

func NewContext() *AppContext {
	doc := js.Global().Get("document")
	root := doc.Call("getElementById", "root")

	return &AppContext{
		Current: root,
		Doc:     doc,
		Root:    Widget(root),
	}
}

func Context() *AppContext {
	return context
}

func (c *AppContext) AppendChild(child BaseWidget) {
	c.Current.Call("appendChild", js.Value(child.Widget))
}

func (c *AppContext) ClientHeight() int {
	return c.Doc.Get("documentElement").Get("clientHeight").Int()
}

func (c *AppContext) ClientWidth() int {
	return c.Doc.Get("documentElement").Get("clientWidth").Int()
}

func (c *AppContext) CreateElement(tag string) Widget {
	return Widget(c.Doc.Call("createElement", tag))
}

func (c *AppContext) CreateElementWithOptions(tag string, options map[string]any) Widget {
	return Widget(c.Doc.Call("createElement", tag, options))
}

func (c *AppContext) CreateElementNS(namespace string, tag string) Widget {
	return Widget(c.Doc.Call("createElementNS", namespace, tag))
}

func (c *AppContext) CurrentPath() string {
	return js.Global().Get("window").Get("location").Get("pathname").String()
}

func (c *AppContext) GetElementByID(id string) *Widget {
	element := c.Doc.Call("getElementById", id)
	if element.IsNull() {
		return nil
	}
	w := Widget(element)
	return &w
}

func (c *AppContext) GoBack() {
	js.Global().Get("window").Get("history").Call("back")
}

func (c *AppContext) Navigate(path string) {
	js.Global().Get("window").Get("history").Call("pushState", nil, "", path)
}

func (c *AppContext) OnNavigate(callback func(path string)) {
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

func (c *AppContext) OpenLink(href string) {
	js.Global().Get("window").Call("open", href, "_blank")
}

func (c *AppContext) OnResize(callback func()) {
	js.Global().Get("window").Set("onresize", js.FuncOf(func(this js.Value, args []js.Value) any {
		if callback != nil {
			utils.SafeGo(func() {
				callback()
			})
		}
		return nil
	}))
}

func (c *AppContext) PushState(state any, title string, url string) {
	js.Global().Get("window").Get("history").Call("pushState", state, title, url)
}
