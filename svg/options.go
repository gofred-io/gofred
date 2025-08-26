package svg

type Options func(svg *Svg)

func Style(styleOptions ...StyleOptions) Options {
	return func(svg *Svg) {
		for _, styleOption := range styleOptions {
			styleOption(svg)
		}
	}
}
