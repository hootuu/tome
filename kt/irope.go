package kt

type IRope interface {
	GetChain() Chain
	GetHeight() Height
	GetTail() KID
}
