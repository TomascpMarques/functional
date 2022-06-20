package elements

type IMG struct {
	Content    ContentList
	Attributes map[string]string
}

func NewIMG(content ...interface{}) *IMG {
	return &IMG{content, make(map[string]string)}
}

func (img *IMG) PushNewElement(e Element) Element {
	img.Content = append(img.Content, e)
	return img
}

func (img IMG) MarkItUp() string {
	return MarkItUpHelper("img", (*[]interface{})(&img.Content), &img.Attributes)
}

func (img *IMG) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		img.Attributes[k] = v
	}
	return img
}

func (img *IMG) ReplaceContent(new Element, pos uint) Element {
	img.Content[pos] = new
	return img
}
