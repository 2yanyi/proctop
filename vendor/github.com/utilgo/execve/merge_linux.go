package execve

import "bytes"

func merge(stdout, stderr *bytes.Buffer) string {
	if stderr.Len() != 0 {
		stdout.WriteString("\n")
		stdout.Write(stderr.Bytes())
	}
	return stdout.String()
}
