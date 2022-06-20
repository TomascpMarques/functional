package elements

type EMBED struct {
	Content    ContentList
	Attributes map[string]string
}

func NewEMBED(content ...interface{}) *EMBED {
	return &EMBED{content, make(map[string]string)}
}

func (embed *EMBED) PushNewElement(e Element) Element {
	embed.Content = append(embed.Content, e)
	return embed
}

func (embed EMBED) MarkItUp() string {
	return MarkItUpHelper("embed", (*[]interface{})(&embed.Content), &embed.Attributes)
}

func (embed *EMBED) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		embed.Attributes[k] = v
	}
	return embed
}

func (embed *EMBED) ReplaceContent(new Element, pos uint) Element {
	embed.Content[pos] = new
	return embed
}
