package elements

type HEAD struct {
	Content    ContentList
	Attributes map[string]string
}

func NewHEAD(content ...interface{}) *HEAD {
	return &HEAD{content, make(map[string]string)}
}

func (head *HEAD) PushNewElement(e Element) Element {
	head.Content = append(head.Content, e)
	return head
}

func (head HEAD) MarkItUp() string {
	return MarkItUpHelper("head", (*[]interface{})(&head.Content), &head.Attributes)
}

func (head *HEAD) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		head.Attributes[k] = v
	}
	return head
}

func (head *HEAD) ReplaceContent(new Element, pos uint) Element {
	head.Content[pos] = new
	return head
}
