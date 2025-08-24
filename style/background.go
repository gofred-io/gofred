package style

import "fmt"

type Background struct {
	Color string
}

func (b *Background) String() string {
	return fmt.Sprintf("background-color: %s;", b.Color)
}
