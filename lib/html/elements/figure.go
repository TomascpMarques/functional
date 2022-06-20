package elements

type FIGURE struct {
	Content    ContentList
	Attributes map[string]string
}

func NewFIGURE(content ...interface{}) *FIGURE {
	return &FIGURE{content, make(map[string]string)}
}

func (figure *FIGURE) PushNewElement(e Element) Element {
	figure.Content = append(figure.Content, e)
	return figure
}

func (figure FIGURE) MarkItUp() string {
	return MarkItUpHelper("figure", (*[]interface{})(&figure.Content), &figure.Attributes)
}

func (figure *FIGURE) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		figure.Attributes[k] = v
	}
	return figure
}

func (figure *FIGURE) ReplaceContent(new Element, pos uint) Element {
	figure.Content[pos] = new
	return figure
}
