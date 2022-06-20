package elements

type FONT struct {
	Content    ContentList
	Attributes map[string]string
}

func NewFONT(content ...interface{}) *FONT {
	return &FONT{content, make(map[string]string)}
}

func (font *FONT) PushNewElement(e Element) Element {
	font.Content = append(font.Content, e)
	return font
}

func (font FONT) MarkItUp() string {
	return MarkItUpHelper("font", (*[]interface{})(&font.Content), &font.Attributes)
}

func (font *FONT) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		font.Attributes[k] = v
	}
	return font
}

func (font *FONT) ReplaceContent(new Element, pos uint) Element {
	font.Content[pos] = new
	return font
}
