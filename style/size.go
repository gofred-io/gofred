package style

import "fmt"

type Size struct {
	Width  int
	Height int
}

func (s *Size) String() string {
	return fmt.Sprintf("width: %dpx; height: %dpx;", s.Width, s.Height)
}
