package path

import (
	. "github.com/gofred-io/gofred/widget"
)

type BPath struct {
	*BaseWidget
}

func Path(data string) *BPath {
	path := &BPath{
		BaseWidget: NewNS("http://www.w3.org/2000/svg", "path"),
	}

	path.d(data)
	return path
}
