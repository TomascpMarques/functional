package elements

type FOOTER struct {
	Content    ContentList
	Attributes map[string]string
}

func NewFOOTER(content ...interface{}) *FOOTER {
	return &FOOTER{content, make(map[string]string)}
}

func (footer *FOOTER) PushNewElement(e Element) Element {
	footer.Content = append(footer.Content, e)
	return footer
}

func (footer FOOTER) MarkItUp() string {
	return MarkItUpHelper("footer", (*[]interface{})(&footer.Content), &footer.Attributes)
}

func (footer *FOOTER) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		footer.Attributes[k] = v
	}
	return footer
}

func (footer *FOOTER) ReplaceContent(new Element, pos uint) Element {
	footer.Content[pos] = new
	return footer
}
