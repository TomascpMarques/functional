package elements

type INPUT struct {
	Content    ContentList
	Attributes map[string]string
}

func NewINPUT(content ...interface{}) *INPUT {
	return &INPUT{content, make(map[string]string)}
}

func (input *INPUT) PushNewElement(e Element) Element {
	input.Content = append(input.Content, e)
	return input
}

func (input INPUT) MarkItUp() string {
	return MarkItUpHelper("input", (*[]interface{})(&input.Content), &input.Attributes)
}

func (input *INPUT) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		input.Attributes[k] = v
	}
	return input
}

func (input *INPUT) ReplaceContent(new Element, pos uint) Element {
	input.Content[pos] = new
	return input
}
