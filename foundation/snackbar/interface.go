package snackbar

import "github.com/gofred-io/gofred/widget"

type ISnackbar interface {
	Show(child widget.BaseWidget)
}
