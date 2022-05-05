package execve

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func merge(stdout, stderr *bytes.Buffer) string {
	if stderr.Len() != 0 {
		stdout.WriteString("\n")
		stdout.Write(stderr.Bytes())
	}
	text, _ := simplifiedchinese.GBK.NewDecoder().String(stdout.String())
	return text
}
