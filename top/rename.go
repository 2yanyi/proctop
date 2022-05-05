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
		if strings.HasSuffix(*name, ".exe") {
			*name = strings.TrimSuffix(*name, ".exe")
		}
	}
	if *name == "java" {
		renameJava(name, commandline)
		return
	} else {
		defer func(bak string) {
			if *name != bak {
				*name += " >"
			}
		}(*name)
	}
	if variable.IsWin {
		renameWindows(name)
	}

	switch {

	// Kernel
	case strings.HasPrefix(*name, "kworker/"):
		*name = "Kworker"
	case strings.HasPrefix(*name, "ksoftirqd/"):
		*name = "Ksoftirqd"
	case strings.HasPrefix(*name, "irq/"):
		*name = "Irq"
	case strings.HasPrefix(*name, "migration/"):
		*name = "Migration"
	case strings.HasPrefix(*name, "idle_inject/"):
		*name = "Idle_inject"
	case strings.HasPrefix(*name, "cpuhp/"):
		*name = "Cpuhp"
	case strings.HasPrefix(*name, "scsi_"):
		*name = "SCSI"
	case strings.HasPrefix(*name, "rcu"):
		*name = "RCU"
	case strings.HasPrefix(*name, "card0-"):
		*name = "card0"

	// System
	case strings.HasPrefix(*name, "upstart"):
		*name = "UPStart"
	case strings.HasPrefix(*name, "indicator"):
		*name = "Indicator"
	case strings.HasPrefix(*name, "systemd"), *name == "(sd-pam)":
		*name = "Systemd"
	case strings.HasPrefix(*name, "dbus"):
		*name = "Dbus"
	case strings.HasPrefix(*name, "ibus"):
		*name = "Ibus"
	case strings.HasPrefix(*name, "cups"):
		*name = "Cups"
	case strings.HasPrefix(*name, "xdg"):
		*name = "Xdg"
	case strings.HasPrefix(*name, "fcitx"):
		*name = "Fcitx"
	case strings.HasPrefix(*name, "evolution"):
		*name = "Evolution"
	case strings.HasPrefix(*name, "pipewire"):
		*name = "PipeWire"
	case strings.HasPrefix(*name, "unity"):
		*name = "Unity-tools"

	// GNOME
	case strings.HasPrefix(*name, "tracker"),
		strings.HasPrefix(*name, "gvfs"),
		strings.HasPrefix(*name, "gdm"),
		strings.HasPrefix(*name, "gsd"),
		strings.HasPrefix(*name, "goa"),
		strings.HasPrefix(*name, "at-spi"),
		strings.HasPrefix(*name, "gnome"), *name == "gjs", *name == "dconf-service":
		switch *name {
		case "gnome-terminal":
			*name = "Terminal"
		case "gnome-disks":
			*name = "disks"
		default:
			*name = "GNOME"
		}

	// case strings.HasPrefix(*name, "tracker"):
	// 	*name = "Tracker"
	// case strings.HasPrefix(*name, "gvfs"):
	// 	*name = "Gvfs"
	// case strings.HasPrefix(*name, "gdm"):
	// 	*name = "Gdm"
	// case strings.HasPrefix(*name, "gsd"):
	// 	*name = "Gsd"
	// case strings.HasPrefix(*name, "goa"):
	// 	*name = "Goa"
	// case strings.HasPrefix(*name, "at-spi"):
	// 	*name = "AT-spi"
	// case strings.HasPrefix(*name, "gnome"):
	// 	{
	// 		switch *name {
	// 		case "gnome-terminal", "gnome-terminal.real":
	// 			*name = "Terminal"
	// 		case "gnome-disks":
	// 			*name = "disks"
	// 		default:
	// 			*name = "GNOME"
	// 		}
	// 	}

	// Database
	case strings.HasPrefix(*name, "clickhouse"):
		*name = "ClickHouse"
	case strings.HasPrefix(*name, "mongo"):
		*name = "MongoDB"
	case strings.HasPrefix(*name, "mysql"):
		*name = "MySQL"
	case strings.HasPrefix(*name, "redis"):
		*name = "Redis"

	// Applications
	case strings.HasPrefix(*name, "chrome"):
		if strings.Contains(*commandline, "chromium") {
			*name = "chromium"
		} else {
			*name = "Google Chrome"
		}
	case strings.HasPrefix(*name, "sysproxy-cmd"):
		*name = "Lantern"
	case strings.HasPrefix(*name, "docker"):
		*name = "Docker"
	case strings.HasPrefix(*name, "PM2"):
		*name = "PM2"
	case strings.HasPrefix(*name, "fsnotifier"):
		*name = "FsNotifier"
	case strings.HasPrefix(*name, "virt"), *name == "libvirtd":
		*name = "Virt-manager"
	case strings.HasPrefix(*name, "BaiduNetdisk"):
		*name = "BaiduDrive"
	case *name == "steamwebhelper":
		*name = "Steam"

	// VM
	case strings.HasPrefix(*name, "VBox"), *name == "VirtualBoxVM":
		if *name != "VBoxClient" {
			*name = "VirtualBoxVM"
		}
	case *name == "vmtoolsd", *name == "vmware-vmblock-fuse", *name == "vmhgfs-fuse", *name == "VGAuthService":
		*name = "VMware-tools"

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
	switch *name {
	case "StartMenuExperienceHost",
		"ApplicationFrameHost",
		"ShellExperienceHost",
		"backgroundTaskHost",
		"TextInputHost",
		"SearchHost",
		"conhost", "svchost", "sihost", "dllhost", "taskhostw",
		"SystemSettings", "SystemSettingsBroker", "Widgets", "LockApp",
		"ctfmon":
		*name = "System Host"
	case "hvsirdpclient", "hvsirpcd", "hvsimgr":
		*name = "Windows Defender"
	case "WindowsTerminal", "OpenConsole":
		*name = "Windows Terminal"
	case "GameBar", "GameBarFTServer", "Video.UI":
		*name = "Xbox GameBar"
	case "HxTsr":
		*name = "Microsoft Office"
	case "WeChatStore", "WsaClient":
		*name = "WeChat"
	case "IObitUninstaler", "UninstallMonitor":
		*name = "IObit Uninstaller"
	case "RtkAudUService64", "RtkUWP":
		*name = "Realtek Audio"
	case "喜马拉雅":
		*name = "Ximalaya"
	}
	switch {
	case strings.HasPrefix(*name, "AMD"), *name == "RadeonSoftware", *name == "cncmd":
		*name = "AMD Radeon Software"
	case strings.HasPrefix(*name, "Armoury"), *name == "Aura Wallpaper Service":
		*name = "ROG ArmouryCrate"
	case strings.HasPrefix(strings.ToLower(*name), "asus"), *name == "AcPowerNotification", *name == "AsHotplugCtrl":
		*name = "ASUS Software"
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
		*name = javaTag + "Hadoop"
		return
	}

	// Java: hbase
	if strings.Contains(*commandline, "-Dhbase.log") {
		*name = javaTag + "Hbase"
		return
	}

	// Java: zookeeper
	if strings.Contains(*commandline, "-Dzookeeper.log") {
		*name = javaTag + "Zookeeper"
		return
	}

	// Java: kafka
	if strings.Contains(*commandline, "-Dkafka.log") {
		*name = javaTag + "Kafka"
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
