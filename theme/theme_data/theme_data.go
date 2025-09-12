package theme_data

type ThemeData struct {
	themeName string
	textTheme textTheme
}

type ThemeDataOption func(themeData *ThemeData)

func NewThemeData(opts ...ThemeDataOption) ThemeData {
	themeData := ThemeData{}

	for _, opt := range opts {
		opt(&themeData)
	}

	return themeData
}

func WithTextTheme(textTheme textTheme) ThemeDataOption {
	return func(themeData *ThemeData) {
		themeData.textTheme = textTheme
	}
}

func WithThemeName(themeName string) ThemeDataOption {
	return func(themeData *ThemeData) {
		themeData.themeName = themeName
	}
}

func (t *ThemeData) TextTheme() textTheme {
	return t.textTheme
}

func (t *ThemeData) Name() string {
	return t.themeName
}
