package elements

type HTML struct {
	Content    ContentList
	Attributes map[string]string
}

func NewHTML(content ...interface{}) *HTML {
	return &HTML{content, make(map[string]string)}
}

func (html *HTML) PushNewElement(e Element) Element {
	html.Content = append(html.Content, e)
	return html
}

func (html HTML) MarkItUp() string {
	return MarkItUpHelper("html", (*[]interface{})(&html.Content), &html.Attributes)
}

func (html *HTML) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		html.Attributes[k] = v
	}
	return html
}

func (html *HTML) ReplaceContent(new Element, pos uint) Element {
	html.Content[pos] = new
	return html
}
