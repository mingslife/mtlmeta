package mtl

type MTLDocument struct {
	children []MTLElement
}

func NewMTLDocument() *MTLDocument {
	return &MTLDocument{}
}

func (elem *MTLDocument) Children() []MTLElement {
	return elem.children
}

func (elem *MTLDocument) AppendChild(child MTLElement) bool {
	elem.children = append(elem.children, child)
	return true
}

func (elem *MTLDocument) String() string {
	str := ""
	for _, child := range elem.children {
		str += child.String()
	}
	return str
}

func (elem *MTLDocument) Bytes() []byte {
	bytes := []byte{}
	for _, child := range elem.children {
		bytes = append(bytes, child.Bytes()...)
	}
	return bytes
}

var _ MTLElement = (*MTLDocument)(nil)
