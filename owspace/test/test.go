package main

import (
	"fmt"
	"r/owspace"
	"time"
)

func main() {
	owspace.New(func(w *owspace.Writer) {
		for i := 0; i <= 100; i++ {
			w.Write(fmt.Sprintf("Downloading... (%d/%d)\n", i, 100))
			time.Sleep(time.Millisecond * 10)
		}
	})
}

/*

 Num Count  Memory                             Name    CPU%  / Core*4
--------------------------------------------------------------------------------------
 13)   1    47.5M                        bamfdaemon     0.0
 16)   1    43.8M                   update-notifier     0.0
 21)   1    39.9M                        notify-osd     0.0
 22)   1   768.0K                         gpg-agent     0.0
 26)   1    16.3M                     zeitgeist-fts     0.0
 27)   1    15.6M                 zeitgeist-datahub     0.0
 30)   1    38.2M                         nm-applet     0.0
 42)   1    12.1M                          whoopsie     0.0
 45)   1     8.6M               window-stack-bridge     0.0
 46)   2    14.0M                           lightdm     0.0
 48)   1    11.1M                  zeitgeist-daemon     0.0
 50)   1    10.6M                  deja-dup-monitor     0.0
 53)   1     4.0M                           dnsmasq     0.0
 57)   1     1.8M                            agetty     0.0
 58)   1     3.6M                          dhclient     0.0
 59)   1     5.3M                          gconfd-2     0.0
 12)   1    46.6M                       hud-service     0.0

*/
