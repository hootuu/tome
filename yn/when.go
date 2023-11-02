package yn

import (
	"fmt"
	"time"
)

type When int64

func NewWhen() When {
	c := time.Now().Unix()
	return When(c)
}

func (w When) S() string {
	return fmt.Sprintf("%d", w)
}

func (w When) Time() time.Time {
	return time.Unix(int64(w), 0)
}
