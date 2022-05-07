package internal

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"runtime"
	"strings"
)

type ProcessorExtensionInformation struct {
	HertzMax string
	Power    string
	Score    string
}

func cpuTitle() (perf, processor string) {
	stat, _ := cpu.Info()
	if len(stat) == 0 {
		return
	}
	switch {
	case strings.HasPrefix(stat[0].ModelName, "AMD"):
		processor = strings.TrimSpace(strings.Split(stat[0].ModelName, "with")[0])
	case strings.HasPrefix(stat[0].ModelName, "Intel"):
		processor = strings.TrimSpace(strings.Split(stat[0].ModelName, "@")[0])
	}
	pei := database[processor]
	if pei.Power != "" {
		perf = fmt.Sprintf("UP%s Hertz(Max:%s Power:%s).T%d", pei.Score, pei.HertzMax, pei.Power, runtime.NumCPU())
	} else {
		perf = fmt.Sprintf("Hertz(%.2fG).T%d", stat[0].Mhz/1000, runtime.NumCPU())
	}
	info, _ := mem.VirtualMemory()
	if info != nil {
		perf += fmt.Sprintf(" - Memory(%s)", SizeFormat(float64(info.Total)))
	}
	return
}

var database = make(map[string]ProcessorExtensionInformation)

var _ = func() error {
	for i := range carefullySelectedCPUs {
		values := strings.Split(carefullySelectedCPUs[i], ",")
		if len(values) != 4 {
			continue
		}
		database[strings.TrimSpace(values[0])] = ProcessorExtensionInformation{
			strings.TrimSpace(values[1]),
			strings.TrimSpace(values[2]),
			strings.TrimSpace(values[3]),
		}
	}
	return nil
}()

var carefullySelectedCPUs = []string{

	// 锐龙™ 线程撕裂者™

	"AMD Ryzen Threadripper PRO 5995WX, 4.50G, 280W, ?",
	"AMD Ryzen Threadripper PRO 5975WX, 4.50G, 280W, ?",
	"AMD Ryzen Threadripper PRO 5965WX, 4.50G, 280W, ?",
	"AMD Ryzen Threadripper PRO 5955WX, 4.50G, 280W, ?",
	"AMD Ryzen Threadripper PRO 5945WX, 4.50G, 280W, ?",
	"AMD Ryzen Threadripper PRO 3995WX, 4.20G, 280W, 1527045",
	"AMD Ryzen Threadripper PRO 3975WX, 4.20G, 280W, 1081952",
	"AMD Ryzen Threadripper PRO 3955WX, 4.30G, 280W, 728395",
	"AMD Ryzen Threadripper 3990X,      4.30G, 280W, 1482287",
	"AMD Ryzen Threadripper 3970X,      4.50G, 280W, 1127921",
	"AMD Ryzen Threadripper 3960X,      4.50G, 280W, 944004",

	// 锐龙™ 6000

	"AMD Ryzen 9 6980HX, 5.00G, 45W, 642142",
	"AMD Ryzen 9 6980HS, 5.00G, 35W, 636496",
	"AMD Ryzen 9 6900HX, 4.90G, 45W, 642144",
	"AMD Ryzen 9 6900HS, 4.90G, 35W, 633897",
	"AMD Ryzen 7 6800H,  4.70G, 45W, 640091",
	"AMD Ryzen 7 6800HS, 4.70G, 35W, 631298",
	"AMD Ryzen 5 6600H,  4.50G, 45W, 628421",
	"AMD Ryzen 5 6600HS, 4.50G, 35W, 618699",

	// 锐龙™ 5000

	"AMD Ryzen 9 5980HX, 4.80G, 45W, 559653",
	"AMD Ryzen 9 5980HS, 4.80G, 35W, 540326",
	"AMD Ryzen 9 5900HX, 4.60G, 45W, 559912",
	"AMD Ryzen 9 5900HS, 4.60G, 35W, 530856",
	"AMD Ryzen 7 5800H,  4.40G, 45W, 539419",
	"AMD Ryzen 7 5800HS, 4.40G, 35W, 487658",
	"AMD Ryzen 5 5600H,  4.20G, 45W, 457236",
	"AMD Ryzen 5 5600HS, 4.20G, 35W, 438791",

	// 酷睿™ 12 Gen

	"Intel(R) Core(TM) i9-12900HK CPU, 5.00G, 45W, 728372",
	"Intel(R) Core(TM) i9-12900H CPU,  5.00G, 45W, 712820",
	"Intel(R) Core(TM) i7-12800H CPU,  4.80G, 45W, 678769",
	"Intel(R) Core(TM) i7-12700H CPU,  4.70G, 45W, 667170",
	"Intel(R) Core(TM) i7-12650H CPU,  4.70G, 45W, 584704",
	"Intel(R) Core(TM) i5-12600H CPU,  4.50G, 45W, 582765",
	"Intel(R) Core(TM) i5-12500H CPU,  4.50G, 45W, 582671",
	"Intel(R) Core(TM) i5-12450H CPU,  4.40G, 45W, 467982",
}
