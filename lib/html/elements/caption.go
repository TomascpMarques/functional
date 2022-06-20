package elements

type CAPTION struct {
	Content    ContentList
	Attributes map[string]string
}

func NewCAPTION(content ...interface{}) *CAPTION {
	return &CAPTION{content, make(map[string]string)}
}

func (caption *CAPTION) PushNewElement(e Element) Element {
	caption.Content = append(caption.Content, e)
	return caption
}

func (caption CAPTION) MarkItUp() string {
	return MarkItUpHelper("caption", (*[]interface{})(&caption.Content), &caption.Attributes)
}

func (caption *CAPTION) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		caption.Attributes[k] = v
	}
	return caption
}

func (caption *CAPTION) ReplaceContent(new Element, pos uint) Element {
	caption.Content[pos] = new
	return caption
}
