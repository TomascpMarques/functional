package elements

type BDO struct {
	Content    ContentList
	Attributes map[string]string
}

func NewBDO(content ...interface{}) *BDO {
	return &BDO{content, make(map[string]string)}
}

func (bdo *BDO) PushNewElement(e Element) Element {
	bdo.Content = append(bdo.Content, e)
	return bdo
}

func (bdo BDO) MarkItUp() string {
	return MarkItUpHelper("bdo", (*[]interface{})(&bdo.Content), &bdo.Attributes)
}

func (bdo *BDO) SetAttributes(attr map[string]string) Element {
	for k, v := range attr {
		bdo.Attributes[k] = v
	}
	return bdo
}

func (bdo *BDO) ReplaceContent(new Element, pos uint) Element {
	bdo.Content[pos] = new
	return bdo
}
