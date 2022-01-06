package homepage

import (
	"fmt"
	"r/scanner"
	"runtime"
)

var cpuMax = float64(runtime.NumCPU() * 100)

func WebsiteMatch(s string, cpu *float64) (_ string) {
	if s == scanner.StatisticsTag {
		return fmt.Sprintf("%.2f%%", *cpu/cpuMax*100)
	}
	if microsoft[s] {
		return "\u001B[0;35;48mmicrosoft\u001B[0m"
	}
	return components[s]
}

var components = map[string]string{
	"msedge.exe": "https://microsoft.com/edge",
	"code.exe":   "https://code.visualstudio.com",
}

var microsoft = map[string]bool{
	"sihost.exe":                  true,
	"startmenuexperiencehost.exe": true,
	"msteams.exe":                 true,
	"openconsole.exe":             true,
	"powershell.exe":              true,
	"hxtsr.exe":                   true,
	"searchhost.exe":              true,
	"onedrive.exe":                true,
	"svchost.exe":                 true,
	"dllhost.exe":                 true,
	"taskhostw.exe":               true,
	"securityhealthsystray.exe":   true,
	"backgroundtaskhost.exe":      true,
	"shellexperiencehost.exe":     true,
	"explorer.exe":                true,
	"windowsterminal.exe":         true,
	"runtimebroker.exe":           true,
	"msteamsupdate.exe":           true,
	"msedgewebview2.exe":          true,
	"smartscreen.exe":             true,
	"microsoft.photos.exe":        true,
	"applicationframehost.exe":    true,
	"winstore.app.exe":            true,
	"minisearchhost.exe":          true,
	"widgets.exe":                 true,
}
