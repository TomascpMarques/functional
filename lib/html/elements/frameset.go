package elements

type FRAMESET struct {
	Content    ContentList
	Attributes map[string]string
}

func NewFRAMESET(content ...interface{}) *FRAMESET {
	return &FRAMESET{content, make(map[string]string)}
}

func (frameset *FRAMESET) PushNewElement(e Element) Element {
	frameset.Content = append(frameset.Content, e)
	return frameset
}

func (frameset FRAMESET) MarkItUp() string {
	return MarkItUpHelper("frameset", (*[]interface{})(&frameset.Content), &frameset.Attributes)
}

func (frameset *FRAMESET) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		frameset.Attributes[k] = v
	}
	return frameset
}

func (frameset *FRAMESET) ReplaceContent(new Element, pos uint) Element {
	frameset.Content[pos] = new
	return frameset
}
