package elements

type H3 struct {
	Content    ContentList
	Attributes map[string]string
}

func NewH3(content ...interface{}) *H3 {
	return &H3{content, make(map[string]string)}
}

func (h3 *H3) PushNewElement(e Element) Element {
	h3.Content = append(h3.Content, e)
	return h3
}

func (h3 H3) MarkItUp() string {
	return MarkItUpHelper("h3", (*[]interface{})(&h3.Content), &h3.Attributes)
}

func (h3 *H3) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		h3.Attributes[k] = v
	}
	return h3
}

func (h3 *H3) ReplaceContent(new Element, pos uint) Element {
	h3.Content[pos] = new
	return h3
}
