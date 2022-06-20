package elements

type HR struct {
	Content    ContentList
	Attributes map[string]string
}

func NewHR(content ...interface{}) *HR {
	return &HR{content, make(map[string]string)}
}

func (hr *HR) PushNewElement(e Element) Element {
	hr.Content = append(hr.Content, e)
	return hr
}

func (hr HR) MarkItUp() string {
	return MarkItUpHelper("hr", (*[]interface{})(&hr.Content), &hr.Attributes)
}

func (hr *HR) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		hr.Attributes[k] = v
	}
	return hr
}

func (hr *HR) ReplaceContent(new Element, pos uint) Element {
	hr.Content[pos] = new
	return hr
}
