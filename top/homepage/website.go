package homepage

var Components = map[string]string{
	"proctop": "https://github.com/matsuwin/proctop",

	// FreeDesktop Projects
	"modemmanager":          "https://github.com/freedesktop/modemmanager",
	"dbus":                  "https://github.com/freedesktop/dbus",
	"upowerd":               "https://github.com/freedesktop/upower",
	"polkitd":               "https://github.com/freedesktop/polkit",
	"accounts-daemon":       "https://github.com/freedesktop/accountsservice",
	"geoclue":               "https://github.com/freedesktop/geoclue",
	"colord":                "https://github.com/freedesktop/colord",
	"pulseaudio":            "https://github.com/freedesktop/pulseaudio",
	"plymouthd":             "https://github.com/freedesktop/plymouth",
	"plymouth":              "https://github.com/freedesktop/plymouth",
	"networkmanager":        "https://github.com/freedesktop/networkmanager",
	"nm-connection-editor":  "https://github.com/freedesktop/NetworkManager",
	"nm-dispatcher":         "https://github.com/freedesktop/NetworkManager/tree/master/src/nm-dispatcher",
	"iio-sensor-proxy":      "https://github.com/hadess/iio-sensor-proxy",
	"switcheroo-control":    "https://github.com/hadess/switcheroo-control",
	"xorg":                  "https://x.org",
	"xwayland":              "https://wayland.freedesktop.org",
	"power-profiles-daemon": "https://gitlab.freedesktop.org/hadess/power-profiles-daemon",
	"lightdm":               "https://github.com/canonical/lightdm",

	// GNOME Components
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
	"simple-scan":   "https://github.com/gnome/simple-scan",
	"gnome":         "https://gnome.org",
	"terminal":      "https://wiki.gnome.org/Apps/Terminal",
	"disks":         "https://wiki.gnome.org/Apps/Disks",
	"nautilus":      "https://wiki.gnome.org/Apps/Files",
	"boxes":         "https://wiki.gnome.org/Apps/Boxes",
	"dconf-editor":  "https://wiki.gnome.org/Apps/DconfEditor",
	"evince":        "https://wiki.gnome.org/Apps/Evince",
	"totem":         "https://wiki.gnome.org/Apps/Videos",
	"seahorse":      "https://wiki.gnome.org/Apps/Seahorse",
	"gedit":         "https://wiki.gnome.org/Apps/Gedit",
	"baobab":        "https://wiki.gnome.org/Apps/Baobab",
	"yelp":          "https://wiki.gnome.org/Apps/Yelp",
	"geary":         "https://wiki.gnome.org/Apps/Geary",
	"gucharmap":     "https://wiki.gnome.org/Apps/Gucharmap",
	"eog":           "https://wiki.gnome.org/Apps/EyeOfGnome",

	// Pop!_OS
	"system76-power":     "https://github.com/pop-os/system76-power",
	"system76-scheduler": "https://github.com/pop-os/system76-scheduler",
	"pop-upgrade":        "https://github.com/pop-os/upgrade",
	"hidpi-daemon":       "https://github.com/pop-os/hidpi-daemon",
	"popsicle-gtk":       "https://github.com/pop-os/popsicle",

	// LXDE Components
	"lxsession":   "https://github.com/lxde/lxsession",
	"lxpanel":     "https://github.com/lxde/lxpanel",
	"lxpolkit":    "https://github.com/lxqt/lxqt-policykit",
	"pcmanfm":     "https://wiki.lxde.org/zh/PCManFM",
	"menu-cached": "https://github.com/lxde/menu-cache",

	// System Components
	"kernel worker":           "kernel",
	"migration":               "kernel",
	"cron":                    "https://github.com/cronie-crond",
	"bluetoothd":              "https://github.com/bluez",
	"hciattach":               "https://github.com/bluez/bluez/blob/master/tools/hciattach.c",
	"obexd":                   "https://github.com/heinervdm/obexd",
	"packagekitd":             "https://github.com/packagekit",
	"ibus":                    "https://github.com/ibus",
	"fusermount":              "https://github.com/libfuse",
	"auditd":                  "https://github.com/linux-audit",
	"touchegg":                "https://github.com/joseexposito/touchegg",
	"thermald":                "https://github.com/intel/thermal_daemon",
	"xdg":                     "https://github.com/flatpak/xdg-desktop-portal",
	"irqbalance":              "https://github.com/irqbalance",
	"fsnotifier":              "https://github.com/jetbrains/intellij-community/tree/master/native/fsnotifier",
	"avahi-daemon":            "https://github.com/lathiat/avahi",
	"rtkit-daemon":            "https://github.com/heftig/rtkit",
	"wpa_supplicant":          "https://github.com/digsrc/wpa_supplicant",
	"udisksd":                 "https://github.com/storaged-project/udisks",
	"cups":                    "https://github.com/openprinting/cups-filters",
	"compiz":                  "https://github.com/compiz-reloaded/compiz",
	"indicator":               "https://github.com/AyatanaIndicators/libayatana-indicator",
	"acpid":                   "https://sourceforge.net/projects/acpid2",
	"bash":                    "https://gnu.org/software/bash",
	"systemd":                 "https://systemd.io",
	"nacl_helper":             "https://nacl.cr.yp.to",
	"io.elementary.appcenter": "https://appcenter.elementary.io",
	"fcitx":                   "https://fcitx-im.org",
	"fwupd":                   "https://fwupd.org",
	"rsyslogd":                "https://rsyslog.com",
	"pipewire":                "https://pipewire.org",
	"unity-tools":             "https://unityx.org",
	"upstart":                 "https://upstart.ubuntu.com",
	"rinetd":                  "https://github.com/samhocevar/rinetd",
	"notify-osd":              "https://launchpad.net/notify-osd",
	"mutter":                  "https://github.com/collects/mutter",
	"dhcpcd":                  "https://github.com/NetworkConfiguration/dhcpcd",

	// Applications
	"firefox":                  "https://firefox.com",
	"geckomain":                "https://wiki.mozilla.org/Gecko",
	"webextensions":            "https://wiki.mozilla.org/WebExtensions",
	"virtualboxvm":             "https://virtualbox.org",
	"chromium":                 "https://chromium.org",
	"google chrome":            "https://google.com/chrome",
	"code":                     "https://code.visualstudio.com",
	"lantern":                  "https://lantern.io",
	"dbeaver":                  "https://dbeaver.io",
	"insomnia":                 "https://insomnia.rest",
	"blender":                  "https://blender.org",
	"kazam":                    "https://launchpad.net/kazam",
	"vlc":                      "https://videolan.org",
	"remmina":                  "https://remmina.org",
	"transmission-gtk":         "https://transmissionbt.com",
	"rpi-imager":               "https://github.com/raspberrypi/rpi-imager",
	"jcef_helper":              "https://github.com/chromiumembedded/java-cef",
	"com.github.donadigo.eddy": "https://github.com/donadigo/eddy",
	"vmware-tools":             "https://github.com/vmware/open-vm-tools",
	"steam":                    "https://store.steampowered.com",
	"powershell":               "https://github.com/powershell",
	"windows terminal":         "https://github.com/microsoft/terminal",
	"amd radeon software":      "https://amd.com/zh-hans/technologies/software",
	"asus software":            "https://www.asus.com/asus-software",
	"rog armourycrate":         "https://rog.asus.com/armoury-crate",
	"xbox gamebar":             "https://microsoft.com/zh-cn/p/xbox-game-bar/9nzkpstsnw4p",
	"dingtalk":                 "https://dingtalk.com",
	"adrive":                   "https://aliyundrive.com",
	"baidudrive":               "https://pan.baidu.com",
	"idman":                    "https://internetdownloadmanager.com",
	"ximalaya":                 "https://ximalaya.com",
	"vmplayer":                 "https://vmware.com/products/workstation-player",
	"vmware-vmx":               "https://vmware.com",
	"clangd":                   "https://clangd.llvm.org",
	"qtwebengineprocess":       "https://ultimatepp.org",
	"realtek audio":            "https://realtek.com",
	"windows defender":         "https://microsoft.com/zh-cn/security",
	"notepads":                 "https://notepadsapp.com",
	"iobit uninstaller":        "https://www.iobit.com/en/advanceduninstaller",
	"wechat":                   "https://weixin.qq.com",
	"taskmgr":                  "任务管理器",

	// Server Components
	"sshd":         "https://openssh.com",
	"node":         "https://nodejs.org",
	"python3":      "https://python.org",
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
	"java":            "https://openjdk.java.net",
	"j/androidstudio": "https://developer.android.com/studio",
	"j/idea":          "https://jetbrains.com/idea",
	"j/goland":        "https://jetbrains.com/go",
	"j/webstorm":      "https://jetbrains.com/webstorm",
	"j/phpstorm":      "https://jetbrains.com/phpstorm",
	"j/ruby":          "https://jetbrains.com/ruby",
	"j/pycharmcore":   "https://jetbrains.com/pycharm",
	"j/rider":         "https://jetbrains.com/rider",
	"j/clion":         "https://jetbrains.com/clion",
	"j/datagrip":      "https://jetbrains.com/datagrip",
	"j/hadoop":        "https://hadoop.apache.org",
	"j/hbase":         "https://hbase.apache.org",
	"j/zookeeper":     "https://zookeeper.apache.org",
	"j/kafka":         "https://kafka.apache.org",

	// Tools
	"apt":          "https://wiki.debian.org/Apt",
	"vim":          "https://vim.org",
	"snapd":        "https://snapcraft.io",
	"snap-store":   "https://snapcraft.io/snap-store",
	"pigz":         "https://zlib.net/pigz",
	"ffmpeg":       "https://ffmpeg.org",
	"gotop":        "https://github.com/xxxserxxx/gotop",
	"git":          "https://git-scm.com",
	"synaptic":     "https://github.com/mvo5/synaptic",
	"bpftrace":     "https://github.com/iovisor/bpftrace",
	"wineserver64": "https://winehq.org",
	"chsime":       "微软拼音输入法",

	// Windows Basic
	"system host":            "basic",
	"wscript":                "basic",
	"cmd":                    "basic",
	"audiodg":                "basic",
	"searchui":               "basic",
	"explorer":               "basic",
	"smartscreen":            "basic",
	"runtimebroker":          "basic",
	"systemsettings":         "basic",
	"systemsettingsbroker":   "basic",
	"securityhealthsystray":  "basic",
	"openconsole":            "basic",
	"msedgewebview2":         "basic",
	"spatialaudiolicensesrv": "basic",
	"rundll32":               "basic",
	"checknetisolation":      "网络隔离诊断工具",
}
