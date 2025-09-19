package codeblock

import (
	"time"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/code"
	gfcode "github.com/gofred-io/gofred/basic/code"
	"github.com/gofred-io/gofred/basic/pre"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/options/spacing"
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

	copied, setCopied := hooks.UseState(false)

	return container.New(
		column.New(
			[]application.BaseWidget{
				row.New(
					[]application.BaseWidget{
						spacer.New(),
						listenable.Builder(copied, func() application.BaseWidget {
							if copied.Value() {
								return container.New(
									icon.New(
										icondata.CheckAll,
										icon.Fill("#2B799B"),
										icon.Width(breakpoint.All(20)),
										icon.Height(breakpoint.All(20)),
									),
									container.BackgroundColor("transparent"),
									container.Padding(breakpoint.All(spacing.All(4))),
								)
							}
							return iconbutton.New(
								icondata.ContentCopy,
								iconbutton.Class("gf-code-block-copy-button"),
								iconbutton.Tooltip("Copy to clipboard"),
								iconbutton.OnClick(func(this application.BaseWidget, e application.Event) {
									utils.CopyToClipboard(innerCode)

									setCopied(true)
									time.AfterFunc(2*time.Second, func() {
										setCopied(false)
									})

									if codeBlock.onCopied != nil {
										codeBlock.onCopied(innerCode)
									}
								}),
							)
						}),
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
