package elements

type H5 struct {
	Content    ContentList
	Attributes map[string]string
}

func NewH5(content ...interface{}) *H5 {
	return &H5{content, make(map[string]string)}
}

func (h5 *H5) PushNewElement(e Element) Element {
	h5.Content = append(h5.Content, e)
	return h5
}

func (h5 H5) MarkItUp() string {
	return MarkItUpHelper("h5", (*[]interface{})(&h5.Content), &h5.Attributes)
}

func (h5 *H5) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		h5.Attributes[k] = v
	}
	return h5
}

func (h5 *H5) ReplaceContent(new Element, pos uint) Element {
	h5.Content[pos] = new
	return h5
}
