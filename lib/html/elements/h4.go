package elements

type H4 struct {
	Content    ContentList
	Attributes map[string]string
}

func NewH4(content ...interface{}) *H4 {
	return &H4{content, make(map[string]string)}
}

func (h4 *H4) PushNewElement(e Element) Element {
	h4.Content = append(h4.Content, e)
	return h4
}

func (h4 H4) MarkItUp() string {
	return MarkItUpHelper("h4", (*[]interface{})(&h4.Content), &h4.Attributes)
}

func (h4 *H4) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		h4.Attributes[k] = v
	}
	return h4
}

func (h4 *H4) ReplaceContent(new Element, pos uint) Element {
	h4.Content[pos] = new
	return h4
}
