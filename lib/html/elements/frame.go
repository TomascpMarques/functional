package elements

type FRAME struct {
	Content    ContentList
	Attributes map[string]string
}

func NewFRAME(content ...interface{}) *FRAME {
	return &FRAME{content, make(map[string]string)}
}

func (frame *FRAME) PushNewElement(e Element) Element {
	frame.Content = append(frame.Content, e)
	return frame
}

func (frame FRAME) MarkItUp() string {
	return MarkItUpHelper("frame", (*[]interface{})(&frame.Content), &frame.Attributes)
}

func (frame *FRAME) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		frame.Attributes[k] = v
	}
	return frame
}

func (frame *FRAME) ReplaceContent(new Element, pos uint) Element {
	frame.Content[pos] = new
	return frame
}
