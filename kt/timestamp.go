package kt

import (
	"fmt"
	"time"
)

type Timestamp uint64

func NewTimestamp() Timestamp {
	return Timestamp(time.Now().Unix())
}

func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}

func (t Timestamp) S() string {
	return fmt.Sprintf("%d", t)
}
