package elements

type AREA struct {
	Content    ContentList
	Attributes map[string]string
}

func NewAREA(content ...interface{}) *AREA {
	return &AREA{content, make(map[string]string)}
}

func (area *AREA) PushNewElement(e Element) Element {
	area.Content = append(area.Content, e)
	return area
}

func (area AREA) MarkItUp() string {
	return MarkItUpHelper("area", (*[]interface{})(&area.Content), &area.Attributes)
}

func (area *AREA) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		area.Attributes[k] = v
	}
	return area
}

func (area *AREA) ReplaceContent(new Element, pos uint) Element {
	area.Content[pos] = new
	return area
}
