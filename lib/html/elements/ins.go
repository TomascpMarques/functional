package elements

type INS struct {
	Content    ContentList
	Attributes map[string]string
}

func NewINS(content ...interface{}) *INS {
	return &INS{content, make(map[string]string)}
}

func (ins *INS) PushNewElement(e Element) Element {
	ins.Content = append(ins.Content, e)
	return ins
}

func (ins INS) MarkItUp() string {
	return MarkItUpHelper("ins", (*[]interface{})(&ins.Content), &ins.Attributes)
}

func (ins *INS) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		ins.Attributes[k] = v
	}
	return ins
}

func (ins *INS) ReplaceContent(new Element, pos uint) Element {
	ins.Content[pos] = new
	return ins
}
