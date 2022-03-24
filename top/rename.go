package top

import (
	"r/data/variable"
	"strings"
)

/**
 * 同名进程合并
 */

func rename(name, commandline *string) {
	if variable.IsWin {
		renameWindows(name)
	}
	if *name == "java" {
		renameJava(name, commandline)
		return
	}
	switch {

	// System
	case strings.HasPrefix(*name, "kworker/"):
		*name = "kworker"
	case strings.HasPrefix(*name, "upstart"):
		*name = "upstart"
	case strings.HasPrefix(*name, "indicator"):
		*name = "indicator"
	case strings.HasPrefix(*name, "systemd"), *name == "(sd-pam)":
		*name = "systemd"
	case strings.HasPrefix(*name, "dbus"):
		*name = "dbus"
	case strings.HasPrefix(*name, "ibus"):
		*name = "ibus"
	case strings.HasPrefix(*name, "cups"):
		*name = "cups"
	case strings.HasPrefix(*name, "xdg"):
		*name = "xdg"
	case strings.HasPrefix(*name, "fcitx"):
		*name = "fcitx"
	case strings.HasPrefix(*name, "evolution"):
		*name = "evolution"
	case strings.HasPrefix(*name, "pipewire"):
		*name = "pipewire"
	case strings.HasPrefix(*name, "unity"):
		*name = "unity-tools"

	// GNOME
	case strings.HasPrefix(*name, "tracker"):
		*name = "tracker"
	case strings.HasPrefix(*name, "gvfs"):
		*name = "gvfs"
	case strings.HasPrefix(*name, "gdm"):
		*name = "gdm"
	case strings.HasPrefix(*name, "gsd"):
		*name = "gsd"
	case strings.HasPrefix(*name, "goa"):
		*name = "goa"
	case strings.HasPrefix(*name, "at-spi"):
		*name = "at-spi"
	case strings.HasPrefix(*name, "gnome"):
		{
			switch *name {
			case "gnome-terminal", "gnome-terminal.real":
				*name = "terminal"
			case "gnome-disks":
				*name = "disks"
			default:
				*name = "gnome-shell"
			}
		}

	// Database
	case strings.HasPrefix(*name, "clickhouse"):
		*name = "clickhouse"
	case strings.HasPrefix(*name, "mongo"):
		*name = "mongodb"
	case strings.HasPrefix(*name, "mysql"):
		*name = "mysql"
	case strings.HasPrefix(*name, "redis"):
		*name = "redis"

	// Applications
	case strings.HasPrefix(*name, "chrome"):
		if strings.Contains(*commandline, "chromium") {
			*name = "chromium"
		} else {
			*name = "chrome"
		}
	case strings.HasPrefix(*name, "sysproxy-cmd"):
		*name = "lantern"
	case strings.HasPrefix(*name, "docker"):
		*name = "docker"
	case strings.HasPrefix(*name, "PM2"):
		*name = "PM2"
	case strings.HasPrefix(*name, "fsnotifier"):
		*name = "fsnotifier"
	case strings.HasPrefix(*name, "virt"), *name == "libvirtd":
		*name = "virt-manager"
	case strings.HasPrefix(*name, "BaiduNetdisk"):
		*name = "BaiduDrive"
	case *name == "steamwebhelper":
		*name = "steam"

	// VM
	case strings.HasPrefix(*name, "VBox"), *name == "VirtualBoxVM":
		if *name != "VBoxClient" {
			*name = "VirtualBoxVM"
		}
	case *name == "vmtoolsd", *name == "vmware-vmblock-fuse", *name == "vmhgfs-fuse", *name == "VGAuthService":
		*name = "vmware-tools"

	// Java
	case *name == "goland64":
		*name = "J/Goland"

	} // END

	// Path -> Name
	if strings.HasPrefix(*name, "/") {
		tmp := *name
		*name = tmp[strings.LastIndex(tmp, "/")+1:]
	}
}

func renameWindows(name *string) {
	if strings.HasSuffix(*name, ".exe") {
		*name = strings.TrimSuffix(*name, ".exe")
	}
	switch *name {
	case "StartMenuExperienceHost",
		"ApplicationFrameHost",
		"ShellExperienceHost",
		"backgroundTaskHost",
		"TextInputHost",
		"SearchHost",
		"conhost", "svchost", "sihost", "dllhost", "taskhostw",
		"SystemSettings", "SystemSettingsBroker",
		"Widgets", "LockApp":
		*name = "SystemHost"
	case "GameBar", "GameBarFTServer", "Video.UI":
		*name = "Xbox GameBar"
	}
	switch {
	case strings.HasPrefix(*name, "AMD"), *name == "RadeonSoftware":
		*name = "AMD Radeon Software"
	case strings.HasPrefix(*name, "Armoury"), *name == "Aura Wallpaper Service":
		*name = "ROG ArmouryCrate"
	case strings.HasPrefix(strings.ToLower(*name), "asus"):
		*name = "ASUS Software"
	}
	if *name == "喜马拉雅" {
		*name = "Ximalaya"
	}
}

func renameJava(name, commandline *string) {

	// Java: JetBrains
	if strings.Contains(*commandline, "-Didea.vendor.name=") {
		for i, value := range strings.Split(*commandline, "-Didea.platform.prefix=") {
			if i == 0 {
				continue
			}
			*name = ""
			for _, char := range value {
				if char == ' ' {
					break
				}
				*name += string(char)
			}
			*name = javaTag + *name
		}
		return
	}

	// Java: hadoop
	if strings.Contains(*commandline, "-Dhadoop.log") {
		*name = javaTag + "hadoop"
		return
	}

	// Java: hbase
	if strings.Contains(*commandline, "-Dhbase.log") {
		*name = javaTag + "hbase"
		return
	}

	// Java: zookeeper
	if strings.Contains(*commandline, "-Dzookeeper.log") {
		*name = javaTag + "zookeeper"
		return
	}

	// Java: kafka
	if strings.Contains(*commandline, "-Dkafka.log") {
		*name = javaTag + "kafka"
		return
	}

	// java: jar
	jars := strings.Count(*commandline, ".jar")
	if jars == 1 {
		values := strings.Split(*commandline, ".jar")[0]
		value := values[strings.LastIndex(values, "/")+1:]
		*name = javaTag + value + ".jar"
		return
	} else if jars > 1 {
		return
	}

	values := strings.Fields(strings.ReplaceAll(*commandline, "=", " "))
	tag := values[len(values)-2]
	value := ""
	if tag[0] == '-' {
		value = values[len(values)-3]
		if value[0] == '/' {
			value = values[len(values)-1]
		}
	} else {
		value = values[len(values)-1]
	}
	*name = javaTag + value
}
