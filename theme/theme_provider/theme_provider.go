package theme_provider

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
)

func New(child application.BaseWidget) application.BaseWidget {
	themeHook, setThemeData := hooks.UseTheme()
	return listenable.Builder(themeHook, func() application.BaseWidget {
		setThemeData(themeHook.ThemeData())
		return child
	})
}
