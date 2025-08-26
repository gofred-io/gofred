package icon

import "github.com/gofred-io/gofred/icondata"

type Options func(icon *Icon)

func Data(data icondata.IconData) Options {
	return func(icon *Icon) {
		icon.Data = data
	}
}
