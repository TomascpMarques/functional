package elements

type ASIDE struct {
	Content    ContentList
	Attributes map[string]string
}

func NewASIDE(content ...interface{}) *ASIDE {
	return &ASIDE{content, make(map[string]string)}
}

func (aside *ASIDE) PushNewElement(e Element) Element {
	aside.Content = append(aside.Content, e)
	return aside
}

func (aside ASIDE) MarkItUp() string {
	return MarkItUpHelper("aside", (*[]interface{})(&aside.Content), &aside.Attributes)
}

func (aside *ASIDE) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		aside.Attributes[k] = v
	}
	return aside
}

func (aside *ASIDE) ReplaceContent(new Element, pos uint) Element {
	aside.Content[pos] = new
	return aside
}
