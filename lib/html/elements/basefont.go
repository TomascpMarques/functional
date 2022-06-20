package elements

type BASEFONT struct {
	Content    ContentList
	Attributes map[string]string
}

func NewBASEFONT(content ...interface{}) *BASEFONT {
	return &BASEFONT{content, make(map[string]string)}
}

func (basefont *BASEFONT) PushNewElement(e Element) Element {
	basefont.Content = append(basefont.Content, e)
	return basefont
}

func (basefont BASEFONT) MarkItUp() string {
	return MarkItUpHelper("basefont", (*[]interface{})(&basefont.Content), &basefont.Attributes)
}

func (basefont *BASEFONT) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		basefont.Attributes[k] = v
	}
	return basefont
}

func (basefont *BASEFONT) ReplaceContent(new Element, pos uint) Element {
	basefont.Content[pos] = new
	return basefont
}
