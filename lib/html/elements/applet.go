package elements

type APPLET struct {
	Content    ContentList
	Attributes map[string]string
}

func NewAPPLET(content ...interface{}) *APPLET {
	return &APPLET{content, make(map[string]string)}
}

func (applet *APPLET) PushNewElement(e Element) Element {
	applet.Content = append(applet.Content, e)
	return applet
}

func (applet APPLET) MarkItUp() string {
	return MarkItUpHelper("applet", (*[]interface{})(&applet.Content), &applet.Attributes)
}

func (applet *APPLET) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		applet.Attributes[k] = v
	}
	return applet
}

func (applet *APPLET) ReplaceContent(new Element, pos uint) Element {
	applet.Content[pos] = new
	return applet
}
