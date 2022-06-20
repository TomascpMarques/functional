package elements

type MAIN struct {
	Content    ContentList
	Attributes map[string]string
}

func NewMAIN(content ...interface{}) *MAIN {
	return &MAIN{content, make(map[string]string)}
}

func (main *MAIN) PushNewElement(e Element) Element {
	main.Content = append(main.Content, e)
	return main
}

func (main MAIN) MarkItUp() string {
	return MarkItUpHelper("main", (*[]interface{})(&main.Content), &main.Attributes)
}

func (main *MAIN) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		main.Attributes[k] = v
	}
	return main
}

func (main *MAIN) ReplaceContent(new Element, pos uint) Element {
	main.Content[pos] = new
	return main
}
