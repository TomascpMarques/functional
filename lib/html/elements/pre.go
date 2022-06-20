package elements

type PRE struct {
	Content    ContentList
	Attributes map[string]string
}

func NewPRE(content ...interface{}) *PRE {
	return &PRE{content, make(map[string]string)}
}

func (pre *PRE) PushNewElement(e Element) Element {
	pre.Content = append(pre.Content, e)
	return pre
}

func (pre PRE) MarkItUp() string {
	return MarkItUpHelper("pre", (*[]interface{})(&pre.Content), &pre.Attributes)
}

func (pre *PRE) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		pre.Attributes[k] = v
	}
	return pre
}

func (pre *PRE) ReplaceContent(new Element, pos uint) Element {
	pre.Content[pos] = new
	return pre
}
