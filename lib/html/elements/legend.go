package elements

type LEGEND struct {
	Content    ContentList
	Attributes map[string]string
}

func NewLEGEND(content ...interface{}) *LEGEND {
	return &LEGEND{content, make(map[string]string)}
}

func (legend *LEGEND) PushNewElement(e Element) Element {
	legend.Content = append(legend.Content, e)
	return legend
}

func (legend LEGEND) MarkItUp() string {
	return MarkItUpHelper("legend", (*[]interface{})(&legend.Content), &legend.Attributes)
}

func (legend *LEGEND) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		legend.Attributes[k] = v
	}
	return legend
}

func (legend *LEGEND) ReplaceContent(new Element, pos uint) Element {
	legend.Content[pos] = new
	return legend
}
