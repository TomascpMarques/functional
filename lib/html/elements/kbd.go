package elements

type KBD struct {
	Content    ContentList
	Attributes map[string]string
}

func NewKBD(content ...interface{}) *KBD {
	return &KBD{content, make(map[string]string)}
}

func (kbd *KBD) PushNewElement(e Element) Element {
	kbd.Content = append(kbd.Content, e)
	return kbd
}

func (kbd KBD) MarkItUp() string {
	return MarkItUpHelper("kbd", (*[]interface{})(&kbd.Content), &kbd.Attributes)
}

func (kbd *KBD) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		kbd.Attributes[k] = v
	}
	return kbd
}

func (kbd *KBD) ReplaceContent(new Element, pos uint) Element {
	kbd.Content[pos] = new
	return kbd
}
