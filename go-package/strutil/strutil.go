package stringutil // <- this is package name

import "strings"

func RemoveDash(s string) string {
	return remove(s, "-")
}

func remove(s string, rmStr string) string {
	return strings.Replace(s, rmStr, "", -1)
}
