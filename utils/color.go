package utils

import (
	"image/color"
	"math"
	"slices"
)

func ColorToHSL(color color.RGBA) (h, s, l float64) {
	R := float64(color.R) / 255.0
	G := float64(color.G) / 255.0
	B := float64(color.B) / 255.0

	cmax := slices.Max([]float64{R, G, B})
	cmin := slices.Min([]float64{R, G, B})
	delta := cmax - cmin

	l = (cmax + cmin) / 2
	if delta == 0 {
		h = 0
		s = 0
	} else {
		switch cmax {
		case R:
			h = math.Mod((G-B)/delta, 6)
		case G:
			h = ((B - R) / delta) + 2
		case B:
			h = ((R - G) / delta) + 4
		}

		h = h * 60
		s = delta / (1 - math.Abs(2*l-1))
	}

	return
}

func HSLToColor(h, s, l float64) color.RGBA {
	var (
		R, G, B uint8
	)

	C := (1 - math.Abs(2*l-1)) * s
	X := C * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - C/2

	if h < 60 {
		R = uint8((C + m) * 255)
		G = uint8((X + m) * 255)
		B = uint8((0 + m) * 255)
	} else if h < 120 {
		R = uint8((X + m) * 255)
		G = uint8((C + m) * 255)
		B = uint8((0 + m) * 255)
	} else if h < 180 {
		R = uint8((0 + m) * 255)
		G = uint8((C + m) * 255)
		B = uint8((X + m) * 255)
	} else if h < 240 {
		R = uint8((0 + m) * 255)
		G = uint8((X + m) * 255)
		B = uint8((C + m) * 255)
	} else if h < 300 {
		R = uint8((X + m) * 255)
		G = uint8((0 + m) * 255)
		B = uint8((C + m) * 255)
	} else {
		R = uint8((C + m) * 255)
		G = uint8((0 + m) * 255)
		B = uint8((X + m) * 255)
	}

	return color.RGBA{
		R: R,
		G: G,
		B: B,
		A: 255,
	}
}

func InvertColor(color color.RGBA) color.RGBA {
	h, s, l := ColorToHSL(color)
	return HSLToColor(h, s, 1-l)
}
