package elements

type H6 struct {
	Content    ContentList
	Attributes map[string]string
}

func NewH6(content ...interface{}) *H6 {
	return &H6{content, make(map[string]string)}
}

func (h6 *H6) PushNewElement(e Element) Element {
	h6.Content = append(h6.Content, e)
	return h6
}

func (h6 H6) MarkItUp() string {
	return MarkItUpHelper("h6", (*[]interface{})(&h6.Content), &h6.Attributes)
}

func (h6 *H6) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		h6.Attributes[k] = v
	}
	return h6
}

func (h6 *H6) ReplaceContent(new Element, pos uint) Element {
	h6.Content[pos] = new
	return h6
}
