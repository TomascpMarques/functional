package elements

type BASE struct {
	Content    ContentList
	Attributes map[string]string
}

func NewBASE(content ...interface{}) *BASE {
	return &BASE{content, make(map[string]string)}
}

func (base *BASE) PushNewElement(e Element) Element {
	base.Content = append(base.Content, e)
	return base
}

func (base BASE) MarkItUp() string {
	return MarkItUpHelper("base", (*[]interface{})(&base.Content), &base.Attributes)
}

func (base *BASE) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		base.Attributes[k] = v
	}
	return base
}

func (base *BASE) ReplaceContent(new Element, pos uint) Element {
	base.Content[pos] = new
	return base
}
