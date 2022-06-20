package elements

type META struct {
	Content    ContentList
	Attributes map[string]string
}

func NewMETA(content ...interface{}) *META {
	return &META{content, make(map[string]string)}
}

func (meta *META) PushNewElement(e Element) Element {
	meta.Content = append(meta.Content, e)
	return meta
}

func (meta META) MarkItUp() string {
	return MarkItUpHelper("meta", (*[]interface{})(&meta.Content), &meta.Attributes)
}

func (meta *META) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		meta.Attributes[k] = v
	}
	return meta
}

func (meta *META) ReplaceContent(new Element, pos uint) Element {
	meta.Content[pos] = new
	return meta
}
