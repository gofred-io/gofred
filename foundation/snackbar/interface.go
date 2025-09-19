package snackbar

import "github.com/gofred-io/gofred/application"

type ISnackbar interface {
	Show(child application.BaseWidget)
}
