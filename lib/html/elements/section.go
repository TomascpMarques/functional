package elements

type SECTION struct {
	Content    ContentList
	Attributes map[string]string
}

func NewSECTION(content ...interface{}) *SECTION {
	return &SECTION{content, make(map[string]string)}
}

func (section *SECTION) PushNewElement(e Element) Element {
	section.Content = append(section.Content, e)
	return section
}

func (section SECTION) MarkItUp() string {
	return MarkItUpHelper("section", (*[]interface{})(&section.Content), &section.Attributes)
}

func (section *SECTION) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		section.Attributes[k] = v
	}
	return section
}

func (section *SECTION) ReplaceContent(new Element, pos uint) Element {
	section.Content[pos] = new
	return section
}
