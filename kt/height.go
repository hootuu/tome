package kt

import "fmt"

type Height uint64

const (
	GenesisHeight Height = 0
)

func (h Height) S() string {
	return fmt.Sprintf("%d", h)
}

func (h Height) Next() Height {
	return h + 1
}

func (h Height) Previous() Height {
	if h <= 0 {
		return GenesisHeight
	}
	return h - 1
}
