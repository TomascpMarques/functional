package elements

type CITE struct {
	Content    ContentList
	Attributes map[string]string
}

func NewCITE(content ...interface{}) *CITE {
	return &CITE{content, make(map[string]string)}
}

func (cite *CITE) PushNewElement(e Element) Element {
	cite.Content = append(cite.Content, e)
	return cite
}

func (cite CITE) MarkItUp() string {
	return MarkItUpHelper("cite", (*[]interface{})(&cite.Content), &cite.Attributes)
}

func (cite *CITE) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		cite.Attributes[k] = v
	}
	return cite
}

func (cite *CITE) ReplaceContent(new Element, pos uint) Element {
	cite.Content[pos] = new
	return cite
}
