package typography

import "strings"

func UcFirst(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}
