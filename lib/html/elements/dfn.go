package elements

type DFN struct {
	Content    ContentList
	Attributes map[string]string
}

func NewDFN(content ...interface{}) *DFN {
	return &DFN{content, make(map[string]string)}
}

func (dfn *DFN) PushNewElement(e Element) Element {
	dfn.Content = append(dfn.Content, e)
	return dfn
}

func (dfn DFN) MarkItUp() string {
	return MarkItUpHelper("dfn", (*[]interface{})(&dfn.Content), &dfn.Attributes)
}

func (dfn *DFN) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		dfn.Attributes[k] = v
	}
	return dfn
}

func (dfn *DFN) ReplaceContent(new Element, pos uint) Element {
	dfn.Content[pos] = new
	return dfn
}
