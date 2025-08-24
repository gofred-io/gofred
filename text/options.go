package text

type Options func(text *text)

func ID(id string) Options {
	return func(t *text) {
		t.SetID(id)
	}
}

func Style(styleOptions ...StyleOptions) Options {
	return func(t *text) {
		for _, styleOption := range styleOptions {
			styleOption(&t.style)
		}
		t.SetStyle(t.style.String())
	}
}

func Text(_text string) Options {
	return func(t *text) {
		t.SetText(_text)
	}
}
