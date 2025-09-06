package spacer

import (
	. "github.com/gofred-io/gofred/basic/div"
)

type FSpacer struct {
	*BDiv
}

func Spacer() *FSpacer {
	s := &FSpacer{
		BDiv: Div(nil),
	}

	s.flex(1)
	return s
}
