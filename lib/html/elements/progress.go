package elements

type PROGRESS struct {
	Content    ContentList
	Attributes map[string]string
}

func NewPROGRESS(content ...interface{}) *PROGRESS {
	return &PROGRESS{content, make(map[string]string)}
}

func (progress *PROGRESS) PushNewElement(e Element) Element {
	progress.Content = append(progress.Content, e)
	return progress
}

func (progress PROGRESS) MarkItUp() string {
	return MarkItUpHelper("progress", (*[]interface{})(&progress.Content), &progress.Attributes)
}

func (progress *PROGRESS) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		progress.Attributes[k] = v
	}
	return progress
}

func (progress *PROGRESS) ReplaceContent(new Element, pos uint) Element {
	progress.Content[pos] = new
	return progress
}
