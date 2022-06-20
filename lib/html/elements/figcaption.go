package elements

type FIGCAPTION struct {
	Content    ContentList
	Attributes map[string]string
}

func NewFIGCAPTION(content ...interface{}) *FIGCAPTION {
	return &FIGCAPTION{content, make(map[string]string)}
}

func (figcaption *FIGCAPTION) PushNewElement(e Element) Element {
	figcaption.Content = append(figcaption.Content, e)
	return figcaption
}

func (figcaption FIGCAPTION) MarkItUp() string {
	return MarkItUpHelper("figcaption", (*[]interface{})(&figcaption.Content), &figcaption.Attributes)
}

func (figcaption *FIGCAPTION) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		figcaption.Attributes[k] = v
	}
	return figcaption
}

func (figcaption *FIGCAPTION) ReplaceContent(new Element, pos uint) Element {
	figcaption.Content[pos] = new
	return figcaption
}
