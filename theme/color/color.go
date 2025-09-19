package color

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/gofred-io/gofred/utils"
)

type Color color.RGBA

func From(value int) Color {
	return Color(color.RGBA{
		R: uint8(value >> 24),
		G: uint8(value >> 16 & 0xFF),
		B: uint8(value >> 8 & 0xFF),
		A: uint8(value & 0xFF),
	})
}

func FromRGBHex(hex string) Color {
	hex = strings.TrimPrefix(hex, "#")

	if len(hex) != 6 && len(hex) != 8 {
		return From(0)
	}

	r, _ := strconv.ParseInt(hex[:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)
	a := int64(255)
	if len(hex) == 8 {
		a, _ = strconv.ParseInt(hex[6:8], 16, 64)
	}

	return Color(color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	})
}

func (c Color) Invert() Color {
	return Color(utils.InvertColor(color.RGBA(c)))
}

func (c Color) String() string {
	rgba := color.RGBA(c)
	return fmt.Sprintf("#%02X%02X%02X%02X", rgba.R, rgba.G, rgba.B, rgba.A)
}
