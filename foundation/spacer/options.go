package spacer

import (
	. "github.com/gofred-io/gofred/breakpoint"
)

func (spacer *FSpacer) flex(flex int) *FSpacer {
	spacer.BDiv.Flex(flex)
	return spacer
}

func (spacer *FSpacer) Height(height int) *FSpacer {
	spacer.BDiv.NoFlex()
	spacer.BDiv.Height(AllBP(height))
	return spacer
}

func (spacer *FSpacer) Width(width int) *FSpacer {
	spacer.BDiv.NoFlex()
	spacer.BDiv.Width(AllBP(width))
	return spacer
}
