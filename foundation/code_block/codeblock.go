package codeblock

import (
	. "github.com/gofred-io/gofred/basic/code"
	. "github.com/gofred-io/gofred/basic/pre"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/icon_button"
	. "github.com/gofred-io/gofred/foundation/row"
	. "github.com/gofred-io/gofred/foundation/spacer"
	. "github.com/gofred-io/gofred/utils"
	. "github.com/gofred-io/gofred/widget"
)

func CodeBlock(code string) Widget {
	return Container(
		Column(
			[]Widget{
				Row(
					[]Widget{
						Spacer(),
						IconButton(ContentCopy).
							Class("gf-code-block-copy-button").
							OnClick(func(this Widget, e Event) {
								CopyToClipboard(code)
							}),
					},
				).Class("gf-code-block-header"),
				Pre(
					Code(code).
						FontSize(14).
						FontColor("#F3F4F6").
						FontWeight("400"),
				).Class("gf-code-block-pre"),
			},
		).Flex(1),
	).Class("gf-code-block")
}
