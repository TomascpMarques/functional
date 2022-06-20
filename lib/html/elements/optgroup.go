package elements

type OPTGROUP struct {
	Content    ContentList
	Attributes map[string]string
}

func NewOPTGROUP(content ...interface{}) *OPTGROUP {
	return &OPTGROUP{content, make(map[string]string)}
}

func (optgroup *OPTGROUP) PushNewElement(e Element) Element {
	optgroup.Content = append(optgroup.Content, e)
	return optgroup
}

func (optgroup OPTGROUP) MarkItUp() string {
	return MarkItUpHelper("optgroup", (*[]interface{})(&optgroup.Content), &optgroup.Attributes)
}

func (optgroup *OPTGROUP) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		optgroup.Attributes[k] = v
	}
	return optgroup
}

func (optgroup *OPTGROUP) ReplaceContent(new Element, pos uint) Element {
	optgroup.Content[pos] = new
	return optgroup
}
