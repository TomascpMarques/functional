package elements

type LINK struct {
	Content    ContentList
	Attributes map[string]string
}

func NewLINK(content ...interface{}) *LINK {
	return &LINK{content, make(map[string]string)}
}

func (link *LINK) PushNewElement(e Element) Element {
	link.Content = append(link.Content, e)
	return link
}

func (link LINK) MarkItUp() string {
	return MarkItUpHelper("link", (*[]interface{})(&link.Content), &link.Attributes)
}

func (link *LINK) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		link.Attributes[k] = v
	}
	return link
}

func (link *LINK) ReplaceContent(new Element, pos uint) Element {
	link.Content[pos] = new
	return link
}
