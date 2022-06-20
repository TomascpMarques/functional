package elements

type DATALIST struct {
	Content    ContentList
	Attributes map[string]string
}

func NewDATALIST(content ...interface{}) *DATALIST {
	return &DATALIST{content, make(map[string]string)}
}

func (datalist *DATALIST) PushNewElement(e Element) Element {
	datalist.Content = append(datalist.Content, e)
	return datalist
}

func (datalist DATALIST) MarkItUp() string {
	return MarkItUpHelper("datalist", (*[]interface{})(&datalist.Content), &datalist.Attributes)
}

func (datalist *DATALIST) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		datalist.Attributes[k] = v
	}
	return datalist
}

func (datalist *DATALIST) ReplaceContent(new Element, pos uint) Element {
	datalist.Content[pos] = new
	return datalist
}
