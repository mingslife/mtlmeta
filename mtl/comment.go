package mtl

type MTLComment struct {
	comment string
}

func NewMTLComment(comment string) *MTLComment {
	return &MTLComment{comment}
}

func (elem *MTLComment) Children() []MTLElement {
	return nil
}

func (elem *MTLComment) AppendChild(child MTLElement) bool {
	return false
}

func (elem *MTLComment) String() string {
	return "//" + elem.comment
}

func (elem *MTLComment) Bytes() []byte {
	return []byte(elem.String())
}

var _ MTLElement = (*MTLComment)(nil)
