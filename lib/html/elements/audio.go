package elements

type AUDIO struct {
	Content    ContentList
	Attributes map[string]string
}

func NewAUDIO(content ...interface{}) *AUDIO {
	return &AUDIO{content, make(map[string]string)}
}

func (audio *AUDIO) PushNewElement(e Element) Element {
	audio.Content = append(audio.Content, e)
	return audio
}

func (audio AUDIO) MarkItUp() string {
	return MarkItUpHelper("audio", (*[]interface{})(&audio.Content), &audio.Attributes)
}

func (audio *AUDIO) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		audio.Attributes[k] = v
	}
	return audio
}

func (audio *AUDIO) ReplaceContent(new Element, pos uint) Element {
	audio.Content[pos] = new
	return audio
}
