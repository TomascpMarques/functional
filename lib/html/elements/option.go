package elements

type OPTION struct {
	Content    ContentList
	Attributes map[string]string
}

func NewOPTION(content ...interface{}) *OPTION {
	return &OPTION{content, make(map[string]string)}
}

func (option *OPTION) PushNewElement(e Element) Element {
	option.Content = append(option.Content, e)
	return option
}

func (option OPTION) MarkItUp() string {
	return MarkItUpHelper("option", (*[]interface{})(&option.Content), &option.Attributes)
}

func (option *OPTION) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		option.Attributes[k] = v
	}
	return option
}

func (option *OPTION) ReplaceContent(new Element, pos uint) Element {
	option.Content[pos] = new
	return option
}
