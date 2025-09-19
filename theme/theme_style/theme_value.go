package theme_style

type themeValue[T any] *T

func ThemeValue[T any](value T) themeValue[T] {
	return &value
}
