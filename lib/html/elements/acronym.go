package elements

type ACRONYM struct {
	Content    ContentList
	Attributes map[string]string
}

func NewACRONYM(content ...interface{}) *ACRONYM {
	return &ACRONYM{content, make(map[string]string)}
}

func (acronym *ACRONYM) PushNewElement(e Element) Element {
	acronym.Content = append(acronym.Content, e)
	return acronym
}

func (acronym ACRONYM) MarkItUp() string {
	return MarkItUpHelper("acronym", (*[]interface{})(&acronym.Content), &acronym.Attributes)
}

func (acronym *ACRONYM) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		acronym.Attributes[k] = v
	}
	return acronym
}

func (acronym *ACRONYM) ReplaceContent(new Element, pos uint) Element {
	acronym.Content[pos] = new
	return acronym
}
