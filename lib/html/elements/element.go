package elements

type ContentList []interface{}

type Element interface {
	MarkItUp() string
	PushNewElement(e Element) Element
	SetAttributes(attr map[string]string) Element
	ReplaceContent(new Element, pos uint) Element
}
