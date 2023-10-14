package ki

type PRI string

func (pri PRI) S() string {
	return string(pri)
}
