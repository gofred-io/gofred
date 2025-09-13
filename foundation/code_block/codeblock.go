package codeblock

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/code"
	gfcode "github.com/gofred-io/gofred/basic/code"
	"github.com/gofred-io/gofred/basic/pre"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/theme"
	"github.com/gofred-io/gofred/theme/theme_style"
	"github.com/gofred-io/gofred/utils"
)

type CodeBlock struct {
	opts     []container.Option
	codeOpts []code.Option
	onCopied func(code string)
}

func New(innerCode string, opts ...Option) application.BaseWidget {
	codeBlock := &CodeBlock{}

	opts = mergeWithDefaultOpts(opts)
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
								utils.CopyToClipboard(innerCode)
								if codeBlock.onCopied != nil {
									codeBlock.onCopied(innerCode)
								}
							}),
						),
					},
					row.Class("gf-code-block-header"),
				),
				pre.New(
					gfcode.New(
						innerCode,
						append(
							[]code.Option{
								code.FontSize(14),
								code.FontColor("#F3F4F6"),
								code.FontWeight("400"),
							},
							codeBlock.codeOpts...,
						)...,
					),
					pre.Class("gf-code-block-pre"),
				),
			},
			column.Flex(1),
			column.Overflow(theme.OverflowTypeHidden),
		),
		codeBlock.opts...,
	)
}

func mergeWithDefaultOpts(opts []Option) []Option {
	return append(
		defaultOpts(),
		opts...,
	)
}

func defaultOpts() []Option {
	return []Option{
		ContainerStyle(defaultThemeData()),
		Class("gf-code-block"),
	}
}

func defaultThemeData() theme_style.ContainerStyle {
	themeHook, _ := hooks.UseTheme()
	themeData := themeHook.ThemeData()
	return themeData.BoxTheme.CodeBlockStyle.Primary
}
