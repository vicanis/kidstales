package parserlib

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func printTree(node *html.Node, indent string) {
	if node == nil {
		return
	}

	if len(node.Data) > 0 {
		sb := strings.Builder{}

		if len(indent) > 0 {
			sb.WriteString(indent)
		}

		sb.WriteString(node.Data)

		if len(node.Attr) > 0 {
			sb.WriteString("[" + printAttrs(node.Attr) + "]")
		}

		fmt.Println(sb.String())
	}

	if node.FirstChild != nil {
		printTree(node.FirstChild, indent+"....")
	}

	if node.NextSibling != nil {
		printTree(node.NextSibling, indent)
	}
}

func printAttrs(attrs []html.Attribute) string {
	pairs := make([]string, len(attrs))

	for i, attr := range attrs {
		pairs[i] = fmt.Sprintf(`%s="%s"`, attr.Key, attr.Val)
	}

	return strings.Join(pairs, " ")
}
