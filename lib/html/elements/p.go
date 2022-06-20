package elements

type ContentList []interface{}

type P struct {
	Content    ContentList
	Attributes map[string]string
}

func NewP(content ...interface{}) *P {
	return &P{content, make(map[string]string)}
}

func (p *P) PushNewElement(e Element) Element {
	p.Content = append(p.Content, e)
	return p
}

func (p P) MarkItUp() string {
	return MarkItUpHelper("p", (*[]interface{})(&p.Content), &p.Attributes)
}

func (p *P) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		p.Attributes[k] = v
	}
	return p
}

func (p *P) ReplaceContent(new Element, pos uint) Element {
	p.Content[pos] = new
	return p
}
