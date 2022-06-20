package elements

type SELECT struct {
	Content    ContentList
	Attributes map[string]string
}

func NewSELECT(content ...interface{}) *SELECT {
	return &SELECT{content, make(map[string]string)}
}

func (Select *SELECT) PushNewElement(e Element) Element {
	Select.Content = append(Select.Content, e)
	return Select
}

func (Select SELECT) MarkItUp() string {
	return MarkItUpHelper("Select", (*[]interface{})(&Select.Content), &Select.Attributes)
}

func (Select *SELECT) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		Select.Attributes[k] = v
	}
	return Select
}

func (Select *SELECT) ReplaceContent(new Element, pos uint) Element {
	Select.Content[pos] = new
	return Select
}
