package code

import (
	. "github.com/gofred-io/gofred/widget"
)

type BCode struct {
	*BaseWidget
}

func Code(innerCode string) *BCode {
	code := &BCode{
		BaseWidget: New("code"),
	}

	code.Set("innerText", innerCode)
	return code
}
