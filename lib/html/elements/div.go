package elements

type Div struct {
	Content    ContentList
	Attributes map[string]string
}

func NewDiv(content ...interface{}) *Div {
	return &Div{content, make(map[string]string)}
}

func (d *Div) PushNewElement(e Element) Element {
	d.Content = append(d.Content, e)
	return d
}

func (d Div) MarkItUp() string {
	return MarkItUpHelper("div", (*[]interface{})(&d.Content), &d.Attributes)
}

func (d *Div) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		d.Attributes[k] = v
	}
	return d
}

func (d *Div) ReplaceContent(new Element, pos uint) Element {
	d.Content[pos] = new
	return d
}
