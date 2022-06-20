package elements

type H2 struct {
	Content    ContentList
	Attributes map[string]string
}

func NewH2(content ...interface{}) *H2 {
	return &H2{content, make(map[string]string)}
}

func (h2 *H2) PushNewElement(e Element) Element {
	h2.Content = append(h2.Content, e)
	return h2
}

func (h2 H2) MarkItUp() string {
	return MarkItUpHelper("h2", (*[]interface{})(&h2.Content), &h2.Attributes)
}

func (h2 *H2) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		h2.Attributes[k] = v
	}
	return h2
}

func (h2 *H2) ReplaceContent(new Element, pos uint) Element {
	h2.Content[pos] = new
	return h2
}
