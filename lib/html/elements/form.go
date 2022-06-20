package elements

type FORM struct {
	Content    ContentList
	Attributes map[string]string
}

func NewFORM(content ...interface{}) *FORM {
	return &FORM{content, make(map[string]string)}
}

func (form *FORM) PushNewElement(e Element) Element {
	form.Content = append(form.Content, e)
	return form
}

func (form FORM) MarkItUp() string {
	return MarkItUpHelper("form", (*[]interface{})(&form.Content), &form.Attributes)
}

func (form *FORM) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		form.Attributes[k] = v
	}
	return form
}

func (form *FORM) ReplaceContent(new Element, pos uint) Element {
	form.Content[pos] = new
	return form
}
