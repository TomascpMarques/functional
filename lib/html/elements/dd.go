package elements

type DD struct {
	Content    ContentList
	Attributes map[string]string
}

func NewDD(content ...interface{}) *DD {
	return &DD{content, make(map[string]string)}
}

func (dd *DD) PushNewElement(e Element) Element {
	dd.Content = append(dd.Content, e)
	return dd
}

func (dd DD) MarkItUp() string {
	return MarkItUpHelper("dd", (*[]interface{})(&dd.Content), &dd.Attributes)
}

func (dd *DD) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		dd.Attributes[k] = v
	}
	return dd
}

func (dd *DD) ReplaceContent(new Element, pos uint) Element {
	dd.Content[pos] = new
	return dd
}
