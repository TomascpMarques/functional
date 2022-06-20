package elements

type METER struct {
	Content    ContentList
	Attributes map[string]string
}

func NewMETER(content ...interface{}) *METER {
	return &METER{content, make(map[string]string)}
}

func (meter *METER) PushNewElement(e Element) Element {
	meter.Content = append(meter.Content, e)
	return meter
}

func (meter METER) MarkItUp() string {
	return MarkItUpHelper("meter", (*[]interface{})(&meter.Content), &meter.Attributes)
}

func (meter *METER) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		meter.Attributes[k] = v
	}
	return meter
}

func (meter *METER) ReplaceContent(new Element, pos uint) Element {
	meter.Content[pos] = new
	return meter
}
