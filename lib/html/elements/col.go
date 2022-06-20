package elements

type COL struct {
	Content    ContentList
	Attributes map[string]string
}

func NewCOL(content ...interface{}) *COL {
	return &COL{content, make(map[string]string)}
}

func (col *COL) PushNewElement(e Element) Element {
	col.Content = append(col.Content, e)
	return col
}

func (col COL) MarkItUp() string {
	return MarkItUpHelper("col", (*[]interface{})(&col.Content), &col.Attributes)
}

func (col *COL) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		col.Attributes[k] = v
	}
	return col
}

func (col *COL) ReplaceContent(new Element, pos uint) Element {
	col.Content[pos] = new
	return col
}
