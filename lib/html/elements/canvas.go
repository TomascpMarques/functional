package elements

type CANVAS struct {
	Content    ContentList
	Attributes map[string]string
}

func NewCANVAS(content ...interface{}) *CANVAS {
	return &CANVAS{content, make(map[string]string)}
}

func (canvas *CANVAS) PushNewElement(e Element) Element {
	canvas.Content = append(canvas.Content, e)
	return canvas
}

func (canvas CANVAS) MarkItUp() string {
	return MarkItUpHelper("canvas", (*[]interface{})(&canvas.Content), &canvas.Attributes)
}

func (canvas *CANVAS) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		canvas.Attributes[k] = v
	}
	return canvas
}

func (canvas *CANVAS) ReplaceContent(new Element, pos uint) Element {
	canvas.Content[pos] = new
	return canvas
}
