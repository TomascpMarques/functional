package elements

type ARTICLE struct {
	Content    ContentList
	Attributes map[string]string
}

func NewARTICLE(content ...interface{}) *ARTICLE {
	return &ARTICLE{content, make(map[string]string)}
}

func (article *ARTICLE) PushNewElement(e Element) Element {
	article.Content = append(article.Content, e)
	return article
}

func (article ARTICLE) MarkItUp() string {
	return MarkItUpHelper("article", (*[]interface{})(&article.Content), &article.Attributes)
}

func (article *ARTICLE) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		article.Attributes[k] = v
	}
	return article
}

func (article *ARTICLE) ReplaceContent(new Element, pos uint) Element {
	article.Content[pos] = new
	return article
}
