package helpers

import (
	"fmt"
	"strings"
)

func SetAttributesHelper[K comparable](src string, target string, attr map[K]K) string {
	return strings.Replace(
		src,
		target,
		func() string {
			var attrb strings.Builder
			for k, v := range attr {
				fmt.Fprintf(&attrb, " %v=\"%v\"", k, v)
			}
			return fmt.Sprintf("%s%s>", target[:len(target)-1], attrb.String())
		}(),
		-1,
	)
}
