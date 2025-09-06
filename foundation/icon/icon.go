package icon

import (
	. "github.com/gofred-io/gofred/basic/path"
	. "github.com/gofred-io/gofred/basic/svg"
	. "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/widget"
)

type FIcon struct {
	*BSvg
}

func Icon(data IconData) *FIcon {
	i := &FIcon{
		BSvg: Svg(
			[]Widget{
				Path(
					string(data),
				),
			},
		),
	}

	i.Class("gf-icon")
	return i
}
