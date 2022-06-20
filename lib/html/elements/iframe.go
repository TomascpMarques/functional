package elements

type IFRAME struct {
	Content    ContentList
	Attributes map[string]string
}

func NewIFRAME(content ...interface{}) *IFRAME {
	return &IFRAME{content, make(map[string]string)}
}

func (iframe *IFRAME) PushNewElement(e Element) Element {
	iframe.Content = append(iframe.Content, e)
	return iframe
}

func (iframe IFRAME) MarkItUp() string {
	return MarkItUpHelper("iframe", (*[]interface{})(&iframe.Content), &iframe.Attributes)
}

func (iframe *IFRAME) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		iframe.Attributes[k] = v
	}
	return iframe
}

func (iframe *IFRAME) ReplaceContent(new Element, pos uint) Element {
	iframe.Content[pos] = new
	return iframe
}
