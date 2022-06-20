package elements

type BLOCKQUOTE struct {
	Content    ContentList
	Attributes map[string]string
}

func NewBLOCKQUOTE(content ...interface{}) *BLOCKQUOTE {
	return &BLOCKQUOTE{content, make(map[string]string)}
}

func (blockquote *BLOCKQUOTE) PushNewElement(e Element) Element {
	blockquote.Content = append(blockquote.Content, e)
	return blockquote
}

func (blockquote BLOCKQUOTE) MarkItUp() string {
	return MarkItUpHelper("blockquote", (*[]interface{})(&blockquote.Content), &blockquote.Attributes)
}

func (blockquote *BLOCKQUOTE) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		blockquote.Attributes[k] = v
	}
	return blockquote
}

func (blockquote *BLOCKQUOTE) ReplaceContent(new Element, pos uint) Element {
	blockquote.Content[pos] = new
	return blockquote
}
