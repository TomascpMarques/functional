package elements

type ABBR struct {
	Content    ContentList
	Attributes map[string]string
}

func NewABBR(content ...interface{}) *ABBR {
	return &ABBR{content, make(map[string]string)}
}

func (abbr *ABBR) PushNewElement(e Element) Element {
	abbr.Content = append(abbr.Content, e)
	return abbr
}

func (abbr ABBR) MarkItUp() string {
	return MarkItUpHelper("abbr", (*[]interface{})(&abbr.Content), &abbr.Attributes)
}

func (abbr *ABBR) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		abbr.Attributes[k] = v
	}
	return abbr
}

func (abbr *ABBR) ReplaceContent(new Element, pos uint) Element {
	abbr.Content[pos] = new
	return abbr
}
