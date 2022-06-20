package elements

type S struct {
	Content    ContentList
	Attributes map[string]string
}

func NewS(content ...interface{}) *S {
	return &S{content, make(map[string]string)}
}

func (s *S) PushNewElement(e Element) Element {
	s.Content = append(s.Content, e)
	return s
}

func (s S) MarkItUp() string {
	return MarkItUpHelper("s", (*[]interface{})(&s.Content), &s.Attributes)
}

func (s *S) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		s.Attributes[k] = v
	}
	return s
}

func (s *S) ReplaceContent(new Element, pos uint) Element {
	s.Content[pos] = new
	return s
}
