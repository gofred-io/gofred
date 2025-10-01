package application

import "syscall/js"

type ClientRect js.Value

func (c ClientRect) Bottom() int {
	return js.Value(c).Get("bottom").Int()
}

func (c ClientRect) Top() int {
	return js.Value(c).Get("top").Int()
}

func (c ClientRect) Left() int {
	return js.Value(c).Get("left").Int()
}

func (c ClientRect) Right() int {
	return js.Value(c).Get("right").Int()
}

func (c ClientRect) Width() int {
	return js.Value(c).Get("width").Int()
}

func (c ClientRect) Height() int {
	return js.Value(c).Get("height").Int()
}

func (c ClientRect) X() int {
	return js.Value(c).Get("x").Int()
}

func (c ClientRect) Y() int {
	return js.Value(c).Get("y").Int()
}
