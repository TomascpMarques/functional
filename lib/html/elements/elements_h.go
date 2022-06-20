package elements

import (
	"fmt"
	"strings"

	"github.com/tomascpmarques/frontline/lib/html/elements/helpers"
)

func MarkItUpHelper[K comparable](
	tagName string,
	srcContent *[]interface{},
	distAttributes *map[K]K,
) string {
	result := fmt.Sprintf("<%s>%s</%s>", tagName, func() string {
		var temp strings.Builder
		for _, value := range *srcContent {
			if trueVal, isElement := value.(Element); isElement {
				temp.WriteString(trueVal.MarkItUp())
				continue
			}
			fmt.Fprintf(&temp, "%v", value)
		}
		return temp.String()
	}(), tagName)

	result = helpers.SetAttributesHelper(result, "<"+tagName+">", *distAttributes)
	return result
}
