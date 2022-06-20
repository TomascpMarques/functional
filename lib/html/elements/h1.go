package elements

type H1 struct {
	Content    ContentList
	Attributes map[string]string
}

func NewH1(content ...interface{}) *H1 {
	return &H1{content, make(map[string]string)}
}

func (h1 *H1) PushNewElement(e Element) Element {
	h1.Content = append(h1.Content, e)
	return h1
}

func (h1 H1) MarkItUp() string {
	return MarkItUpHelper("h1", (*[]interface{})(&h1.Content), &h1.Attributes)
}

func (h1 *H1) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		h1.Attributes[k] = v
	}
	return h1
}

func (h1 *H1) ReplaceContent(new Element, pos uint) Element {
	h1.Content[pos] = new
	return h1
}
