package elements

type DEL struct {
	Content    ContentList
	Attributes map[string]string
}

func NewDEL(content ...interface{}) *DEL {
	return &DEL{content, make(map[string]string)}
}

func (del *DEL) PushNewElement(e Element) Element {
	del.Content = append(del.Content, e)
	return del
}

func (del DEL) MarkItUp() string {
	return MarkItUpHelper("del", (*[]interface{})(&del.Content), &del.Attributes)
}

func (del *DEL) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		del.Attributes[k] = v
	}
	return del
}

func (del *DEL) ReplaceContent(new Element, pos uint) Element {
	del.Content[pos] = new
	return del
}
