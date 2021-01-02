package mtl

type MTLElement interface {
	Children() []MTLElement
	AppendChild(MTLElement) bool
	String() string
	Bytes() []byte
}
