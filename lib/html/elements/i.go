package elements

type I struct {
	Content    ContentList
	Attributes map[string]string
}

func NewI(content ...interface{}) *I {
	return &I{content, make(map[string]string)}
}

func (i *I) PushNewElement(e Element) Element {
	i.Content = append(i.Content, e)
	return i
}

func (i I) MarkItUp() string {
	return MarkItUpHelper("i", (*[]interface{})(&i.Content), &i.Attributes)
}

func (i *I) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		i.Attributes[k] = v
	}
	return i
}

func (i *I) ReplaceContent(new Element, pos uint) Element {
	i.Content[pos] = new
	return i
}
