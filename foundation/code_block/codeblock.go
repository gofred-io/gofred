package codeblock

import (
	gfcode "github.com/gofred-io/gofred/basic/code"
	"github.com/gofred-io/gofred/basic/pre"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/utils"
	"github.com/gofred-io/gofred/widget"
)

func New(code string) widget.Widget {
	return container.New(
		column.New(
			[]widget.Widget{
				row.New(
					[]widget.Widget{
						spacer.New(),
						iconbutton.New(
							icondata.ContentCopy,
							iconbutton.Class("gf-code-block-copy-button"),
							iconbutton.OnClick(func(this widget.Widget, e widget.Event) {
								utils.CopyToClipboard(code)
							}),
						),
					},
					row.Class("gf-code-block-header"),
				),
				pre.New(
					gfcode.New(
						code,
						gfcode.FontSize(14),
						gfcode.FontColor("#F3F4F6"),
						gfcode.FontWeight("400"),
					),
					pre.Class("gf-code-block-pre"),
				),
			},
			column.Flex(1),
		),
		container.Class("gf-code-block"),
	)
}
