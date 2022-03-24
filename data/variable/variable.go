package variable

import "runtime"

var IsWin = func() bool {
	return runtime.GOOS == "windows"
}()
