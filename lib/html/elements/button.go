package elements

type BUTTON struct {
	Content    ContentList
	Attributes map[string]string
}

func NewBUTTON(content ...interface{}) *BUTTON {
	return &BUTTON{content, make(map[string]string)}
}

func (button *BUTTON) PushNewElement(e Element) Element {
	button.Content = append(button.Content, e)
	return button
}

func (button BUTTON) MarkItUp() string {
	return MarkItUpHelper("button", (*[]interface{})(&button.Content), &button.Attributes)
}

func (button *BUTTON) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		button.Attributes[k] = v
	}
	return button
}

func (button *BUTTON) ReplaceContent(new Element, pos uint) Element {
	button.Content[pos] = new
	return button
}
