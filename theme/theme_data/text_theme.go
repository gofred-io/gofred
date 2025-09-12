package theme_data

type textTheme struct {
	codeBlockStyle TextStyleCollection
	textStyle      TextStyleCollection
}

type TextThemeOption func(textTheme *textTheme)

func WithCodeBlockStyle(codeBlockStyle TextStyleCollection) TextThemeOption {
	return func(textTheme *textTheme) {
		textTheme.codeBlockStyle = codeBlockStyle
	}
}

func WithTextStyle(textStyle TextStyleCollection) TextThemeOption {
	return func(textTheme *textTheme) {
		textTheme.textStyle = textStyle
	}
}

func TextTheme(opts ...TextThemeOption) textTheme {
	textTheme := textTheme{}

	for _, opt := range opts {
		opt(&textTheme)
	}

	return textTheme
}

func (t textTheme) CodeBlockStyle() TextStyleCollection {
	return t.codeBlockStyle
}

func (t textTheme) TextStyle() TextStyleCollection {
	return t.textStyle
}
