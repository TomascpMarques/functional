package elements

type NOSCRIPT struct {
	Content    ContentList
	Attributes map[string]string
}

func NewNOSCRIPT(content ...interface{}) *NOSCRIPT {
	return &NOSCRIPT{content, make(map[string]string)}
}

func (noscript *NOSCRIPT) PushNewElement(e Element) Element {
	noscript.Content = append(noscript.Content, e)
	return noscript
}

func (noscript NOSCRIPT) MarkItUp() string {
	return MarkItUpHelper("noscript", (*[]interface{})(&noscript.Content), &noscript.Attributes)
}

func (noscript *NOSCRIPT) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		noscript.Attributes[k] = v
	}
	return noscript
}

func (noscript *NOSCRIPT) ReplaceContent(new Element, pos uint) Element {
	noscript.Content[pos] = new
	return noscript
}
