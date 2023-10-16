package value

import "time"

type Period string

const (
	NilPeriod   Period = ""
	TotalPeriod Period = "$"
)

func (p Period) S() string {
	return string(p)
}

func DatePeriod(dt time.Time) Period {
	now := time.Now()
	fmtDt := now.Format("20060102")
	return Period("DT_" + fmtDt)
}
