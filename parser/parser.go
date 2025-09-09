package parser

import (
	"fmt"
	"strings"
)

func getTypeSpecifier(varType string) string {
	res := ""
	if varType == "string" {
		res = "%s"
	} else if strings.HasPrefix(varType, "int") {
		res = "%d"
	} else if strings.HasPrefix(varType, "float") {
		res = "%f"
	}

	return res

}
func Parse(format string, a []any) (string, []any) {
	types := []string{}
	for _, param := range a {
		types = append(types, getTypeSpecifier(fmt.Sprintf("%T", param)))
	}

	var sb strings.Builder
	i, j := 0, 0
	for i < len(format) {
		if format[i] != '{' {
			sb.WriteString(string(format[i]))
			i++
			continue
		}

		for i < len(format) && format[i] != '}' {
			i++
		}

		sb.WriteString(types[j])
		j++
		i++
	}

	return sb.String(), a
}
