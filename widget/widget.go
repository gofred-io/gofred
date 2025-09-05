package widget

type Widget interface {
	GetBaseWidget() *BaseWidget
	ReplaceWith(widget Widget)
}
