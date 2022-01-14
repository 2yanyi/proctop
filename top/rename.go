package top

import (
	"r/flagset"
	"strings"
)

// 多个进程合并
func rename(name, commandline *string) {

	if *flagset.Java {
		if *name == "java" {
			renameJava(name, commandline)
			return
		}
	}

	switch {

	// System
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
		*name = "chrome"
	case strings.HasPrefix(*name, "sysproxy-cmd"):
		*name = "lantern"
	case strings.HasPrefix(*name, "docker"):
		*name = "docker"
	case strings.HasPrefix(*name, "PM2"):
		*name = "PM2"
	case strings.HasPrefix(*name, "VBox"), *name == "VirtualBoxVM":
		if *name != "VBoxClient" {
			*name = "VirtualBoxVM"
		}
	case strings.HasPrefix(*name, "virt"), *name == "libvirtd":
		*name = "virt-manager"
	}

	// path > name
	if strings.HasPrefix(*name, "/") {
		tmp := *name
		*name = tmp[strings.LastIndex(tmp, "/")+1:]
	}
}

func renameJava(name, commandline *string) {

	// Java: JetBrains
	if strings.Contains(*commandline, "-Didea.vendor.name=JetBrains") {
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
			*name = javaTag + strings.ToLower(*name)
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
		*name = *commandline
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
