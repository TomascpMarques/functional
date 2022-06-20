package elements

type SAMP struct {
	Content    ContentList
	Attributes map[string]string
}

func NewSAMP(content ...interface{}) *SAMP {
	return &SAMP{content, make(map[string]string)}
}

func (samp *SAMP) PushNewElement(e Element) Element {
	samp.Content = append(samp.Content, e)
	return samp
}

func (samp SAMP) MarkItUp() string {
	return MarkItUpHelper("samp", (*[]interface{})(&samp.Content), &samp.Attributes)
}

func (samp *SAMP) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		samp.Attributes[k] = v
	}
	return samp
}

func (samp *SAMP) ReplaceContent(new Element, pos uint) Element {
	samp.Content[pos] = new
	return samp
}
