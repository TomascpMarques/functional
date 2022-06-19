package elements

import (
	"fmt"
	"strings"

	"github.com/tomascpmarques/frontline/lib/html/elements/helpers"
)

type ContentList []interface{}

type P struct {
	Content    ContentList
	Attributes map[string]string
}

func NewP(content ...interface{}) *P {
	return &P{content, make(map[string]string)}
}

func (p *P) PushNewElement(e Element) {
	p.Content = append(p.Content, e)
}

func (p P) MarkItUp() string {
	result := fmt.Sprintf("<p>%s</p>", func() string {
		var temp strings.Builder
		for _, value := range p.Content {
			if trueVal, isElement := value.(Element); isElement {
				temp.WriteString(trueVal.MarkItUp())
				continue
			}
			fmt.Fprintf(&temp, "%v", value)
		}
		return temp.String()
	}())

	result = helpers.SetAttributesHelper(result, `<p>`, p.Attributes)
	return result
}

func (p *P) SetAttributes(attr map[string]string) {
	for k, v := range attr {
		p.Attributes[k] = v
	}
}

func (p *P) ReplaceContent(new Element, pos uint) {
	p.Content[pos] = new
}
