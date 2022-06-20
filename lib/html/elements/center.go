package elements

type CENTER struct {
	Content    ContentList
	Attributes map[string]string
}

func NewCENTER(content ...interface{}) *CENTER {
	return &CENTER{content, make(map[string]string)}
}

func (center *CENTER) PushNewElement(e Element) Element {
	center.Content = append(center.Content, e)
	return center
}

func (center CENTER) MarkItUp() string {
	return MarkItUpHelper("center", (*[]interface{})(&center.Content), &center.Attributes)
}

func (center *CENTER) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		center.Attributes[k] = v
	}
	return center
}

func (center *CENTER) ReplaceContent(new Element, pos uint) Element {
	center.Content[pos] = new
	return center
}
