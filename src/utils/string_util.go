package utils

import (
	"bytes"
)

func StrJoin(strs ...string) string {
	var buf bytes.Buffer
	for _, str := range strs {
		buf.WriteString(str)
	}
	return buf.String()
}
