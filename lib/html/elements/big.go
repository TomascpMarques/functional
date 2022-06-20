package elements

type BIG struct {
	Content    ContentList
	Attributes map[string]string
}

func NewBIG(content ...interface{}) *BIG {
	return &BIG{content, make(map[string]string)}
}

func (big *BIG) PushNewElement(e Element) Element {
	big.Content = append(big.Content, e)
	return big
}

func (big BIG) MarkItUp() string {
	return MarkItUpHelper("big", (*[]interface{})(&big.Content), &big.Attributes)
}

func (big *BIG) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		big.Attributes[k] = v
	}
	return big
}

func (big *BIG) ReplaceContent(new Element, pos uint) Element {
	big.Content[pos] = new
	return big
}
