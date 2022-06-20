package elements

type NAV struct {
	Content    ContentList
	Attributes map[string]string
}

func NewNAV(content ...interface{}) *NAV {
	return &NAV{content, make(map[string]string)}
}

func (nav *NAV) PushNewElement(e Element) Element {
	nav.Content = append(nav.Content, e)
	return nav
}

func (nav NAV) MarkItUp() string {
	return MarkItUpHelper("nav", (*[]interface{})(&nav.Content), &nav.Attributes)
}

func (nav *NAV) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		nav.Attributes[k] = v
	}
	return nav
}

func (nav *NAV) ReplaceContent(new Element, pos uint) Element {
	nav.Content[pos] = new
	return nav
}
