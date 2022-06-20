package elements

type ADDRESS struct {
	Content    ContentList
	Attributes map[string]string
}

func NewADDRESS(content ...interface{}) *ADDRESS {
	return &ADDRESS{content, make(map[string]string)}
}

func (address *ADDRESS) PushNewElement(e Element) Element {
	address.Content = append(address.Content, e)
	return address
}

func (address ADDRESS) MarkItUp() string {
	return MarkItUpHelper("address", (*[]interface{})(&address.Content), &address.Attributes)
}

func (address *ADDRESS) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		address.Attributes[k] = v
	}
	return address
}

func (address *ADDRESS) ReplaceContent(new Element, pos uint) Element {
	address.Content[pos] = new
	return address
}
