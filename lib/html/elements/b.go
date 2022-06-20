package elements

type B struct {
	Content    ContentList
	Attributes map[string]string
}

func NewB(content ...interface{}) *B {
	return &B{content, make(map[string]string)}
}

func (b *B) PushNewElement(e Element) Element {
	b.Content = append(b.Content, e)
	return b
}

func (b B) MarkItUp() string {
	return MarkItUpHelper("b", (*[]interface{})(&b.Content), &b.Attributes)
}

func (b *B) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		b.Attributes[k] = v
	}
	return b
}

func (b *B) ReplaceContent(new Element, pos uint) Element {
	b.Content[pos] = new
	return b
}
