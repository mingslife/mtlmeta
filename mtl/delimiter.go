package mtl

type MTLDelimiter struct {
	crlf bool
}

func NewMTLDelimiter(crlf bool) *MTLDelimiter {
	return &MTLDelimiter{crlf}
}

func (elem *MTLDelimiter) Children() []MTLElement {
	return nil
}

func (elem *MTLDelimiter) AppendChild(child MTLElement) bool {
	return false
}

func (elem *MTLDelimiter) String() string {
	if elem.crlf {
		return "\r\n"
	}
	return "\n"
}

func (elem *MTLDelimiter) Bytes() []byte {
	return []byte(elem.String())
}

var _ MTLElement = (*MTLDelimiter)(nil)
