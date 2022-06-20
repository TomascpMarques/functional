package elements

type MAP struct {
	Content    ContentList
	Attributes map[string]string
}

func NewMAP(content ...interface{}) *MAP {
	return &MAP{content, make(map[string]string)}
}

func (Map *MAP) PushNewElement(e Element) Element {
	Map.Content = append(Map.Content, e)
	return Map
}

func (Map MAP) MarkItUp() string {
	return MarkItUpHelper("map", (*[]interface{})(&Map.Content), &Map.Attributes)
}

func (Map *MAP) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		Map.Attributes[k] = v
	}
	return Map
}

func (Map *MAP) ReplaceContent(new Element, pos uint) Element {
	Map.Content[pos] = new
	return Map
}
