package mtl

import "fmt"

type MTLMaterialItem struct {
	key   string
	value string
}

func NewMTLMaterialItem(key, value string) *MTLMaterialItem {
	return &MTLMaterialItem{key, value}
}

func (elem *MTLMaterialItem) Children() []MTLElement {
	return nil
}

func (elem *MTLMaterialItem) AppendChild(child MTLElement) bool {
	return false
}

func (elem *MTLMaterialItem) String() string {
	return fmt.Sprintf("%s %s", elem.key, elem.value)
}

func (elem *MTLMaterialItem) Bytes() []byte {
	return []byte(elem.String())
}

type MTLMaterial struct {
	name     string
	children []MTLElement
	body     string
}

func NewMTLMaterial(name, body string) *MTLMaterial {
	return &MTLMaterial{name, nil, body}
}

func (elem *MTLMaterial) Children() []MTLElement {
	// return elem.children
	return nil
}

func (elem *MTLMaterial) AppendChild(child MTLElement) bool {
	// switch child.(type) {
	// case *MTLDelimiter, *MTLMaterialItem:
	// 	elem.children = append(elem.children, child)
	// 	return true
	// default:
	// 	return false
	// }
	return false
}

func (elem *MTLMaterial) String() string {
	// str := fmt.Sprintf("#define material \"%s\" {", elem.name)
	// for _, child := range elem.children {
	// 	str += child.String()
	// }
	// str += "}"
	// return str
	str := fmt.Sprintf("#define material \"%s\" ", elem.name)
	str += elem.body
	return str
}

func (elem *MTLMaterial) Bytes() []byte {
	// bytes := []byte(fmt.Sprintf("#define material \"%s\" {", elem.name))
	// for _, child := range elem.children {
	// 	bytes = append(bytes, child.Bytes()...)
	// }
	// bytes = append(bytes, '}')
	// return bytes
	return []byte(elem.String())
}

func (elem *MTLMaterial) Name() string {
	return elem.name
}

func (elem *MTLMaterial) SetName(name string) {
	elem.name = name
}

var _ MTLElement = (*MTLMaterialItem)(nil)
var _ MTLElement = (*MTLMaterial)(nil)
