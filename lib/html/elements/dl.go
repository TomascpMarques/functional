package elements

type DL struct {
	Content    ContentList
	Attributes map[string]string
}

func NewDL(content ...interface{}) *DL {
	return &DL{content, make(map[string]string)}
}

func (dl *DL) PushNewElement(e Element) Element {
	dl.Content = append(dl.Content, e)
	return dl
}

func (dl DL) MarkItUp() string {
	return MarkItUpHelper("dl", (*[]interface{})(&dl.Content), &dl.Attributes)
}

func (dl *DL) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		dl.Attributes[k] = v
	}
	return dl
}

func (dl *DL) ReplaceContent(new Element, pos uint) Element {
	dl.Content[pos] = new
	return dl
}
