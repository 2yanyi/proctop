package top

import (
	"fmt"
	"r/scanner"
	"runtime"
	"strings"
)

var cpuMax = float64(runtime.NumCPU() * 100)

func descriptionMatch(s string, cpu *float64) (_ string) {
	if s == scanner.StatisticsTag {
		return fmt.Sprintf("%.2f%%", *cpu/cpuMax)
	}
	if coreutils[s] {
		return "\u001B[0;37;48mcoreutils\u001B[0m"
	}
	if utilLinux[s] {
		return "\u001B[0;37;48mutil-linux\u001B[0m"
	}
	if strings.HasPrefix(components[s], "https://github.com") {
		return strings.Join([]string{"\u001B[0;35;48m", components[s]}, "")
	}
	return components[s]
}

var components = map[string]string{

	// FreeDesktop Projects
	"modemmanager":       "https://github.com/freedesktop/modemmanager",
	"dbus":               "https://github.com/freedesktop/dbus",
	"upowerd":            "https://github.com/freedesktop/upower",
	"polkitd":            "https://github.com/freedesktop/polkit",
	"accounts-daemon":    "https://github.com/freedesktop/accountsservice",
	"geoclue":            "https://github.com/freedesktop/geoclue",
	"colord":             "https://github.com/freedesktop/colord",
	"pulseaudio":         "https://github.com/freedesktop/pulseaudio",
	"plymouthd":          "https://github.com/freedesktop/plymouth",
	"plymouth":           "https://github.com/freedesktop/plymouth",
	"networkmanager":     "https://github.com/freedesktop/networkmanager",
	"nm-dispatcher":      "https://github.com/freedesktop/NetworkManager/tree/master/src/nm-dispatcher",
	"iio-sensor-proxy":   "https://github.com/hadess/iio-sensor-proxy",
	"switcheroo-control": "https://github.com/hadess/switcheroo-control",

	// GNOME Components Source
	"file-roller":   "https://github.com/gnome/file-roller",
	"gvfs":          "https://github.com/gnome/gvfs",
	"gsd":           "https://github.com/gnome/gnome-settings-daemon",
	"goa":           "https://github.com/gnome/gnome-online-accounts",
	"gjs":           "https://github.com/gnome/gjs",
	"gdm":           "https://github.com/gnome/gdm",
	"at-spi":        "https://github.com/gnome/at-spi2-core",
	"evolution":     "https://github.com/gnome/evolution",
	"dconf-service": "https://github.com/gnome/dconf",
	"tracker":       "https://github.com/gnome/tracker",

	// GNOME Components
	"gnome-shell":  "https://gnome.org",
	"terminal":     "https://wiki.gnome.org/Apps/Terminal",
	"disks":        "https://wiki.gnome.org/Apps/Disks",
	"nautilus":     "https://wiki.gnome.org/Apps/Files",
	"boxes":        "https://wiki.gnome.org/Apps/Boxes",
	"dconf-editor": "https://wiki.gnome.org/Apps/DconfEditor",
	"evince":       "https://wiki.gnome.org/Apps/Evince",

	// System Components Source
	"system76-power": "https://github.com/pop-os/system76-power",
	"pop-upgrade":    "https://github.com/pop-os/upgrade",
	"hidpi-daemon":   "https://github.com/pop-os/hidpi-daemon",
	"cron":           "https://github.com/cronie-crond",
	"bluetoothd":     "https://github.com/bluez",
	"packagekitd":    "https://github.com/packagekit",
	"ibus":           "https://github.com/ibus",
	"fusermount":     "https://github.com/libfuse",
	"auditd":         "https://github.com/linux-audit",
	"touchegg":       "https://github.com/joseexposito/touchegg",
	"thermald":       "https://github.com/intel/thermal_daemon",
	"xdg":            "https://github.com/flatpak/xdg-desktop-portal",
	"irqbalance":     "https://github.com/irqbalance",
	"fsnotifier":     "https://github.com/jetbrains/intellij-community/tree/master/native/fsnotifier",
	"avahi-daemon":   "https://github.com/lathiat/avahi",
	"rtkit-daemon":   "https://github.com/heftig/rtkit",
	"wpa_supplicant": "https://github.com/digsrc/wpa_supplicant",
	"udisksd":        "https://github.com/storaged-project/udisks",
	"cups":           "https://github.com/openprinting/cups-filters",
	"compiz":         "https://github.com/compiz-reloaded/compiz",
	"indicator":      "https://github.com/AyatanaIndicators/libayatana-indicator",

	// System Components Website
	"bash":                    "https://gnu.org/software/bash",
	"systemd":                 "https://systemd.io",
	"xorg":                    "https://x.org",
	"nacl_helper":             "https://nacl.cr.yp.to",
	"io.elementary.appcenter": "https://appcenter.elementary.io",
	"fcitx":                   "https://fcitx-im.org",
	"fwupd":                   "https://fwupd.org",
	"rsyslogd":                "https://rsyslog.com",
	"pipewire":                "https://pipewire.org",
	"unity-tools":             "https://unityx.org",
	"upstart":                 "https://upstart.ubuntu.com",
	"rinetd":                  "https://github.com/samhocevar/rinetd",

	// Applications
	"virtualbox": "https://virtualbox.org",
	"chrome":     "https://google.com/chrome",
	"firefox":    "https://firefox.com",
	"code":       "https://code.visualstudio.com",
	"lantern":    "https://lantern.io",
	"dbeaver":    "https://dbeaver.io",
	"insomnia":   "https://insomnia.rest",
	"blender":    "https://blender.org",
	"kazam":      "https://launchpad.net/kazam",
	"vlc":        "https://videolan.org",

	// Server Components
	"sshd":         "https://openssh.com",
	"node":         "https://nodejs.org",
	"supervisord":  "http://supervisord.org",
	"pm2":          "https://pm2.io",
	"qemu-kvm":     "https://qemu.org",
	"docker":       "https://docker.io",
	"containerd":   "https://containerd.io",
	"clickhouse":   "https://clickhouse.com",
	"mongodb":      "https://mongodb.com",
	"mysql":        "https://mysql.com",
	"redis":        "https://redis.io",
	"minio":        "https://min.io",
	"nginx":        "https://nginx.org",
	"xvnc":         "https://tigervnc.org",
	"virt-manager": "https://virt-manager.org",

	// Java
	"java":           "https://openjdk.java.net",
	"java:idea":      "https://jetbrains.com/idea",
	"java:goland":    "https://jetbrains.com/go",
	"java:hadoop":    "https://hadoop.apache.org",
	"java:hbase":     "https://hbase.apache.org",
	"java:zookeeper": "https://zookeeper.apache.org",
	"java:kafka":     "https://kafka.apache.org",

	// Tools
	"snapd":   "https://snapcraft.io",
	"pigz":    "https://zlib.net/pigz",
	"ffmpeg":  "https://ffmpeg.org",
	"vim":     "https://vim.org",
	"gotop":   "https://github.com/xxxserxxx/gotop",
	"proctop": "https://github.com/matsuwin/proctop",
	"go":      "https://go.dev",
	"python":  "https://python.org",
}
