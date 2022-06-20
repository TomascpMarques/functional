package elements

type OBJECT struct {
	Content    ContentList
	Attributes map[string]string
}

func NewOBJECT(content ...interface{}) *OBJECT {
	return &OBJECT{content, make(map[string]string)}
}

func (object *OBJECT) PushNewElement(e Element) Element {
	object.Content = append(object.Content, e)
	return object
}

func (object OBJECT) MarkItUp() string {
	return MarkItUpHelper("object", (*[]interface{})(&object.Content), &object.Attributes)
}

func (object *OBJECT) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		object.Attributes[k] = v
	}
	return object
}

func (object *OBJECT) ReplaceContent(new Element, pos uint) Element {
	object.Content[pos] = new
	return object
}
