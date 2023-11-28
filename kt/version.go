package kt

import "fmt"

type Version uint64

const (
	DefaultVersion Version = 1
)

func (v Version) S() string {
	return fmt.Sprintf("%d", v)
}
