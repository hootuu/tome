package ki

type Mnemonic string

func (m Mnemonic) S() string {
	return string(m)
}
