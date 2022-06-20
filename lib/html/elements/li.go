package elements

type LI struct {
	Content    ContentList
	Attributes map[string]string
}

func NewLI(content ...interface{}) *LI {
	return &LI{content, make(map[string]string)}
}

func (li *LI) PushNewElement(e Element) Element {
	li.Content = append(li.Content, e)
	return li
}

func (li LI) MarkItUp() string {
	return MarkItUpHelper("li", (*[]interface{})(&li.Content), &li.Attributes)
}

func (li *LI) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		li.Attributes[k] = v
	}
	return li
}

func (li *LI) ReplaceContent(new Element, pos uint) Element {
	li.Content[pos] = new
	return li
}
