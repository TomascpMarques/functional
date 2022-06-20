package elements

type Q struct {
	Content    ContentList
	Attributes map[string]string
}

func NewQ(content ...interface{}) *Q {
	return &Q{content, make(map[string]string)}
}

func (q *Q) PushNewElement(e Element) Element {
	q.Content = append(q.Content, e)
	return q
}

func (q Q) MarkItUp() string {
	return MarkItUpHelper("q", (*[]interface{})(&q.Content), &q.Attributes)
}

func (q *Q) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		q.Attributes[k] = v
	}
	return q
}

func (q *Q) ReplaceContent(new Element, pos uint) Element {
	q.Content[pos] = new
	return q
}
