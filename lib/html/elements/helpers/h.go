package helpers

import (
	"fmt"
	"strings"
)

func SetAttributesHelper(src string, target string, attr map[string]string) string {
	return strings.Replace(
		src,
		target,
		func() string {
			var attrb strings.Builder
			for k, v := range attr {
				fmt.Fprintf(&attrb, " %s=\"%s\"", k, v)
			}
			return fmt.Sprintf("%s%s>", target[:len(target)-1], attrb.String())
		}(),
		-1,
	)
}
