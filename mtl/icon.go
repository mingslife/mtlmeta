package mtl

import "io/ioutil"

type MTLIconKind string

const (
	MTLIconKindBinary = MTLIconKind("binary")
	MTLIconKindString = MTLIconKind("string")
)

var (
	MTLIconMarkBinary = []byte{0x01, 0x02, 0x03, 0x62}
	MTLIconMarkString = []byte{0x62, 0x58, 0x33, 0x30}
)

type MTLIcon struct {
	kind MTLIconKind
	meta []byte
	data []byte
}

func NewMTLIcon(kind MTLIconKind, meta, data []byte) *MTLIcon {
	return &MTLIcon{kind, meta, data}
}

func (elem *MTLIcon) Children() []MTLElement {
	return nil
}

func (elem *MTLIcon) AppendChild(child MTLElement) bool {
	return false
}

func (elem *MTLIcon) String() string {
	return "[MTLIcon:" + string(elem.kind) + "]"
}

func (elem *MTLIcon) Bytes() []byte {
	bytes := []byte("icon ")
	switch elem.kind {
	case MTLIconKindBinary:
		bytes = append(bytes, MTLIconMarkBinary...)
		bytes = append(bytes, elem.meta...)
		bytes = append(bytes, elem.data...)
	case MTLIconKindString:
		bytes = append(bytes, MTLIconMarkString...)
		bytes = append(bytes, elem.data...)
	}
	return bytes
}

func (elem *MTLIcon) Save(path string) {
	ioutil.WriteFile(path, elem.data, 0644)
}

var _ MTLElement = (*MTLIcon)(nil)
