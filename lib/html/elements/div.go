package elements

type DIV struct {
	Content    ContentList
	Attributes map[string]string
}

func NewDIV(content ...interface{}) *DIV {
	return &DIV{content, make(map[string]string)}
}

func (div *DIV) PushNewElement(e Element) Element {
	div.Content = append(div.Content, e)
	return div
}

func (div DIV) MarkItUp() string {
	return MarkItUpHelper("div", (*[]interface{})(&div.Content), &div.Attributes)
}

func (div *DIV) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		div.Attributes[k] = v
	}
	return div
}

func (div *DIV) ReplaceContent(new Element, pos uint) Element {
	div.Content[pos] = new
	return div
}
