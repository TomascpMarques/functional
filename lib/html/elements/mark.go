package elements

type MARK struct {
	Content    ContentList
	Attributes map[string]string
}

func NewMARK(content ...interface{}) *MARK {
	return &MARK{content, make(map[string]string)}
}

func (mark *MARK) PushNewElement(e Element) Element {
	mark.Content = append(mark.Content, e)
	return mark
}

func (mark MARK) MarkItUp() string {
	return MarkItUpHelper("mark", (*[]interface{})(&mark.Content), &mark.Attributes)
}

func (mark *MARK) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		mark.Attributes[k] = v
	}
	return mark
}

func (mark *MARK) ReplaceContent(new Element, pos uint) Element {
	mark.Content[pos] = new
	return mark
}
