package elements

type FIELDSET struct {
	Content    ContentList
	Attributes map[string]string
}

func NewFIELDSET(content ...interface{}) *FIELDSET {
	return &FIELDSET{content, make(map[string]string)}
}

func (fieldset *FIELDSET) PushNewElement(e Element) Element {
	fieldset.Content = append(fieldset.Content, e)
	return fieldset
}

func (fieldset FIELDSET) MarkItUp() string {
	return MarkItUpHelper("fieldset", (*[]interface{})(&fieldset.Content), &fieldset.Attributes)
}

func (fieldset *FIELDSET) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		fieldset.Attributes[k] = v
	}
	return fieldset
}

func (fieldset *FIELDSET) ReplaceContent(new Element, pos uint) Element {
	fieldset.Content[pos] = new
	return fieldset
}
