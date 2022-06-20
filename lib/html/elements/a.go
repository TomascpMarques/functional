package elements

type A struct {
	Content    ContentList
	Attributes map[string]string
}

func NewA(content ...interface{}) *A {
	return &A{content, make(map[string]string)}
}

func (a *A) PushNewElement(e Element) Element {
	a.Content = append(a.Content, e)
	return a
}

func (a A) MarkItUp() string {
	return MarkItUpHelper("a", (*[]interface{})(&a.Content), &a.Attributes)
}

func (a *A) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		a.Attributes[k] = v
	}
	return a
}

func (a *A) ReplaceContent(new Element, pos uint) Element {
	a.Content[pos] = new
	return a
}
