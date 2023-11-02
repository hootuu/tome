package bk

type Rope interface {
	GetChain() Chain
	GetTailKnot() KnotIDX
	GetTailInvariable() BID
}
