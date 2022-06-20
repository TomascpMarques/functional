package elements

type CODE struct {
	Content    ContentList
	Attributes map[string]string
}

func NewCODE(content ...interface{}) *CODE {
	return &CODE{content, make(map[string]string)}
}

func (code *CODE) PushNewElement(e Element) Element {
	code.Content = append(code.Content, e)
	return code
}

func (code CODE) MarkItUp() string {
	return MarkItUpHelper("code", (*[]interface{})(&code.Content), &code.Attributes)
}

func (code *CODE) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		code.Attributes[k] = v
	}
	return code
}

func (code *CODE) ReplaceContent(new Element, pos uint) Element {
	code.Content[pos] = new
	return code
}
