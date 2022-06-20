package elements

type LABEL struct {
	Content    ContentList
	Attributes map[string]string
}

func NewLABEL(content ...interface{}) *LABEL {
	return &LABEL{content, make(map[string]string)}
}

func (label *LABEL) PushNewElement(e Element) Element {
	label.Content = append(label.Content, e)
	return label
}

func (label LABEL) MarkItUp() string {
	return MarkItUpHelper("label", (*[]interface{})(&label.Content), &label.Attributes)
}

func (label *LABEL) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		label.Attributes[k] = v
	}
	return label
}

func (label *LABEL) ReplaceContent(new Element, pos uint) Element {
	label.Content[pos] = new
	return label
}
