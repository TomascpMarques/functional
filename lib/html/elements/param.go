package elements

type PARAM struct {
	Content    ContentList
	Attributes map[string]string
}

func NewPARAM(content ...interface{}) *PARAM {
	return &PARAM{content, make(map[string]string)}
}

func (param *PARAM) PushNewElement(e Element) Element {
	param.Content = append(param.Content, e)
	return param
}

func (param PARAM) MarkItUp() string {
	return MarkItUpHelper("param", (*[]interface{})(&param.Content), &param.Attributes)
}

func (param *PARAM) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		param.Attributes[k] = v
	}
	return param
}

func (param *PARAM) ReplaceContent(new Element, pos uint) Element {
	param.Content[pos] = new
	return param
}
