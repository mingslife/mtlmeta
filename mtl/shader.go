package mtl

import "fmt"

type MTLShaderItem struct {
	key   string
	value string
}

func NewMTLShaderItem(key, value string) *MTLShaderItem {
	return &MTLShaderItem{key, value}
}

func (elem *MTLShaderItem) Children() []MTLElement {
	return nil
}

func (elem *MTLShaderItem) AppendChild(child MTLElement) bool {
	return false
}

func (elem *MTLShaderItem) String() string {
	return fmt.Sprintf("\"%s\" %s,", elem.key, elem.value)
}

func (elem *MTLShaderItem) Bytes() []byte {
	return []byte(elem.String())
}

type MTLShader struct {
	index    int
	name     string
	children []MTLElement
	body     string
}

func NewMTLShader(index int, name, body string) *MTLShader {
	return &MTLShader{index, name, nil, body}
}

func (elem *MTLShader) Children() []MTLElement {
	// return elem.children
	return nil
}

func (elem *MTLShader) AppendChild(child MTLElement) bool {
	// switch child.(type) {
	// case *MTLDelimiter, *MTLShaderItem:
	// 	elem.children = append(elem.children, child)
	// 	return true
	// default:
	// 	return false
	// }
	return false
}

func (elem *MTLShader) String() string {
	// str := fmt.Sprintf("#define shader %d %s {", elem.index, elem.name)
	// for _, child := range elem.children {
	// 	str += child.String()
	// }
	// str += "}"
	// return str
	str := fmt.Sprintf("#define shader %d %s ", elem.index, elem.name)
	str += elem.body
	return str
}

func (elem *MTLShader) Bytes() []byte {
	// bytes := []byte(fmt.Sprintf("#define shader %d %s {", elem.index, elem.name))
	// for _, child := range elem.children {
	// 	bytes = append(bytes, child.Bytes()...)
	// }
	// bytes = append(bytes, '}')
	// return bytes
	return []byte(elem.String())
}

var _ MTLElement = (*MTLShaderItem)(nil)
var _ MTLElement = (*MTLShader)(nil)
