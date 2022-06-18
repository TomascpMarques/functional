package elements

/* Inherit html attributes for each element*/
type Element struct {
	InnerText  string
	InnerHtml  string
	Id         string
	Attributes []string
	Style      string
}

func (e *Element) NewElement(text string, html string, id string, style string, attributes ...string) Element {
	return Element{
		InnerText:  text,
		InnerHtml:  html,
		Id:         id,
		Attributes: attributes,
		Style:      style,
	}
}
