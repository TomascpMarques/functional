package elements

type BODY struct {
	Content    ContentList
	Attributes map[string]string
}

func NewBODY(content ...interface{}) *BODY {
	return &BODY{content, make(map[string]string)}
}

func (body *BODY) PushNewElement(e Element) Element {
	body.Content = append(body.Content, e)
	return body
}

func (body BODY) MarkItUp() string {
	return MarkItUpHelper("body", (*[]interface{})(&body.Content), &body.Attributes)
}

func (body *BODY) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		body.Attributes[k] = v
	}
	return body
}

func (body *BODY) ReplaceContent(new Element, pos uint) Element {
	body.Content[pos] = new
	return body
}
