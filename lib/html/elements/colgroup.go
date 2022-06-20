package elements

type COLGROUP struct {
	Content    ContentList
	Attributes map[string]string
}

func NewCOLGROUP(content ...interface{}) *COLGROUP {
	return &COLGROUP{content, make(map[string]string)}
}

func (colgroup *COLGROUP) PushNewElement(e Element) Element {
	colgroup.Content = append(colgroup.Content, e)
	return colgroup
}

func (colgroup COLGROUP) MarkItUp() string {
	return MarkItUpHelper("colgroup", (*[]interface{})(&colgroup.Content), &colgroup.Attributes)
}

func (colgroup *COLGROUP) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		colgroup.Attributes[k] = v
	}
	return colgroup
}

func (colgroup *COLGROUP) ReplaceContent(new Element, pos uint) Element {
	colgroup.Content[pos] = new
	return colgroup
}
