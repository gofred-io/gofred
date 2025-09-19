package theme_provider

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/foundation/scaffold"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
)

func New(child application.BaseWidget) application.BaseWidget {
	themeHook, setThemeData := hooks.UseTheme()
	return listenable.Builder(themeHook, func() application.BaseWidget {
		setThemeData(themeHook.ThemeData())
		refreshScaffold()
		return child
	})
}

func refreshScaffold() {
	scaffold := scaffold.Get()
	if scaffold == nil {
		return
	}

	scaffold.Refresh()
}
