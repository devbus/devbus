package models

type TextType int8

func (t TextType) String() string {
	switch t {
	case TEXT:
		return "TEXT"
	case HTML:
		return "HTML"
	case MD:
		return "MD"
	default:
		return "UNKNOWN"
	}
}

const (
	TEXT TextType = 0
	HTML TextType = 1
	MD   TextType = 2
)
