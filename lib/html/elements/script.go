package elements

type SCRIPT struct {
	Content    ContentList
	Attributes map[string]string
}

func NewSCRIPT(content ...interface{}) *SCRIPT {
	return &SCRIPT{content, make(map[string]string)}
}

func (script *SCRIPT) PushNewElement(e Element) Element {
	script.Content = append(script.Content, e)
	return script
}

func (script SCRIPT) MarkItUp() string {
	return MarkItUpHelper("script", (*[]interface{})(&script.Content), &script.Attributes)
}

func (script *SCRIPT) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		script.Attributes[k] = v
	}
	return script
}

func (script *SCRIPT) ReplaceContent(new Element, pos uint) Element {
	script.Content[pos] = new
	return script
}
