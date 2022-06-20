package elements

type HEADER struct {
	Content    ContentList
	Attributes map[string]string
}

func NewHEADER(content ...interface{}) *HEADER {
	return &HEADER{content, make(map[string]string)}
}

func (header *HEADER) PushNewElement(e Element) Element {
	header.Content = append(header.Content, e)
	return header
}

func (header HEADER) MarkItUp() string {
	return MarkItUpHelper("header", (*[]interface{})(&header.Content), &header.Attributes)
}

func (header *HEADER) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		header.Attributes[k] = v
	}
	return header
}

func (header *HEADER) ReplaceContent(new Element, pos uint) Element {
	header.Content[pos] = new
	return header
}
