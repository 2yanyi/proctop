package execve

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
)

func Args(dir string, args []string) (_ string) {
	if len(args) != 0 {
		stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
		cmd := exec.Command(args[0])
		cmd.Stdout, cmd.Stderr = stdout, stderr
		cmd.Args = args
		if dir != "" {
			cmd.Dir = filepath.Dir(dir)
		}
		_ = cmd.Run()
		return strings.TrimSpace(merge(stdout, stderr))
	}
	return
}
