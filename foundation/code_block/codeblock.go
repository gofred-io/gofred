package codeblock

import (
	"github.com/gofred-io/gofred/application"
	gfcode "github.com/gofred-io/gofred/basic/code"
	"github.com/gofred-io/gofred/basic/pre"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/theme"
	"github.com/gofred-io/gofred/utils"
)

type CodeBlock struct {
	onCopied func(code string)
}

func New(code string, opts ...Option) application.BaseWidget {
	codeBlock := &CodeBlock{}

	for _, option := range opts {
		option(codeBlock)
	}

	return container.New(
		column.New(
			[]application.BaseWidget{
				row.New(
					[]application.BaseWidget{
						spacer.New(),
						iconbutton.New(
							icondata.ContentCopy,
							iconbutton.Class("gf-code-block-copy-button"),
							iconbutton.Tooltip("Copy to clipboard"),
							iconbutton.OnClick(func(this application.BaseWidget, e application.Event) {
								utils.CopyToClipboard(code)
								if codeBlock.onCopied != nil {
									codeBlock.onCopied(code)
								}
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
			column.Overflow(theme.OverflowTypeHidden),
		),
		container.Class("gf-code-block"),
	)
}
