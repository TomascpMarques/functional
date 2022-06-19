package elements

type Element interface {
	MarkItUp() string
	SetAttributes(attr map[string]string)
	ReplaceContent(new Element, pos uint)
	PushNewElement(e Element)
}
