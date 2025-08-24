package icon

import "github.com/man.go/mango/icondata"

type Options func(icon *Icon)

func Data(data icondata.IconData) Options {
	return func(icon *Icon) {
		icon.Data = data
	}
}
