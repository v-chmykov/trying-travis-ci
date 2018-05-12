package test

import (
	"bytes"
)

// Concat - Concatenate 2 strings
func Concat(str1, str2 string) string {
	var b bytes.Buffer

	b.WriteString(str1)
	b.WriteString(str2)

	return b.String()
}
