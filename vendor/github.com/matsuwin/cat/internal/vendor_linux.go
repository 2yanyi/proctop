package internal

import "strings"

func vendor() (_ string) {
	return strings.TrimSpace(String("/sys/class/dmi/id/sys_vendor"))
}
