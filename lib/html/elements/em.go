package elements

type EM struct {
	Content    ContentList
	Attributes map[string]string
}

func NewEM(content ...interface{}) *EM {
	return &EM{content, make(map[string]string)}
}

func (em *EM) PushNewElement(e Element) Element {
	em.Content = append(em.Content, e)
	return em
}

func (em EM) MarkItUp() string {
	return MarkItUpHelper("em", (*[]interface{})(&em.Content), &em.Attributes)
}

func (em *EM) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		em.Attributes[k] = v
	}
	return em
}

func (em *EM) ReplaceContent(new Element, pos uint) Element {
	em.Content[pos] = new
	return em
}
