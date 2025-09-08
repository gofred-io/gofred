package codeblock

type Option func(codeBlock *CodeBlock)

func OnCopied(onCopied func(code string)) Option {
	return func(codeBlock *CodeBlock) {
		codeBlock.onCopied = onCopied
	}
}
