package elements

type OL struct {
	Content    ContentList
	Attributes map[string]string
}

func NewOL(content ...interface{}) *OL {
	return &OL{content, make(map[string]string)}
}

func (ol *OL) PushNewElement(e Element) Element {
	ol.Content = append(ol.Content, e)
	return ol
}

func (ol OL) MarkItUp() string {
	return MarkItUpHelper("ol", (*[]interface{})(&ol.Content), &ol.Attributes)
}

func (ol *OL) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		ol.Attributes[k] = v
	}
	return ol
}

func (ol *OL) ReplaceContent(new Element, pos uint) Element {
	ol.Content[pos] = new
	return ol
}
