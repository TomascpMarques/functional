package elements

type SMALL struct {
	Content    ContentList
	Attributes map[string]string
}

func NewSMALL(content ...interface{}) *SMALL {
	return &SMALL{content, make(map[string]string)}
}

func (small *SMALL) PushNewElement(e Element) Element {
	small.Content = append(small.Content, e)
	return small
}

func (small SMALL) MarkItUp() string {
	return MarkItUpHelper("small", (*[]interface{})(&small.Content), &small.Attributes)
}

func (small *SMALL) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		small.Attributes[k] = v
	}
	return small
}

func (small *SMALL) ReplaceContent(new Element, pos uint) Element {
	small.Content[pos] = new
	return small
}
