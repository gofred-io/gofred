package codeblock

import (
	"github.com/gofred-io/gofred/basic/code"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/theme/theme_style"
)

type Option func(codeBlock *CodeBlock)

func Class(class string) Option {
	return func(codeBlock *CodeBlock) {
		codeBlock.opts = append(codeBlock.opts, container.Class(class))
	}
}

func ContainerStyle(containerStyle theme_style.ContainerStyle) Option {
	return func(codeBlock *CodeBlock) {
		codeBlock.opts = append(codeBlock.opts, container.ContainerStyle(containerStyle))
	}
}

func OnCopied(onCopied func(code string)) Option {
	return func(codeBlock *CodeBlock) {
		codeBlock.onCopied = onCopied
	}
}

func TextStyle(textStyle theme_style.TextStyle) Option {
	return func(codeBlock *CodeBlock) {
		codeBlock.codeOpts = append(codeBlock.codeOpts, code.TextStyle(textStyle))
	}
}
