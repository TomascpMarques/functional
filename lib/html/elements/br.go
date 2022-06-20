package elements

type BR struct {
	Content    ContentList
	Attributes map[string]string
}

func NewBR(content ...interface{}) *BR {
	return &BR{content, make(map[string]string)}
}

func (br *BR) PushNewElement(e Element) Element {
	br.Content = append(br.Content, e)
	return br
}

func (br BR) MarkItUp() string {
	return MarkItUpHelper("br", (*[]interface{})(&br.Content), &br.Attributes)
}

func (br *BR) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		br.Attributes[k] = v
	}
	return br
}

func (br *BR) ReplaceContent(new Element, pos uint) Element {
	br.Content[pos] = new
	return br
}
