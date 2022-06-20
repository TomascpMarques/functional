package elements

type DT struct {
	Content    ContentList
	Attributes map[string]string
}

func NewDT(content ...interface{}) *DT {
	return &DT{content, make(map[string]string)}
}

func (dt *DT) PushNewElement(e Element) Element {
	dt.Content = append(dt.Content, e)
	return dt
}

func (dt DT) MarkItUp() string {
	return MarkItUpHelper("dt", (*[]interface{})(&dt.Content), &dt.Attributes)
}

func (dt *DT) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		dt.Attributes[k] = v
	}
	return dt
}

func (dt *DT) ReplaceContent(new Element, pos uint) Element {
	dt.Content[pos] = new
	return dt
}
