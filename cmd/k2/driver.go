package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/meadori/k2/pkg/scan/host"
	"github.com/meadori/k2/pkg/scan/port"
)

const usage = `k2 ( https://github/meadori/k2 )
Usage: [Scan Type(s)] {target specification}
`

var doConnectScan bool
var portRangeSpec string

func init() {
	flag.BoolVar(&doConnectScan, "sT", true, "TCP connect scan")
	flag.StringVar(&portRangeSpec, "p", "1-1024", "Scan specified ports")
}

func parseRange(rangeSpec string) (uint, uint) {
	re := regexp.MustCompile(`^(?P<begin>\d+)(?P<dash>-)?(?P<end>\d+)?$`)
	matches := re.FindStringSubmatch(rangeSpec)
	if matches == nil {
		fmt.Println("error: invalid port range specification")
		os.Exit(1)
	}

	begin, err := strconv.ParseUint(matches[1], 10, 32)
	if err != nil {
		fmt.Println("error: invalid port range specification")
		os.Exit(1)
	}

	end := begin
	if len(matches[2]) != 0 {
		if len(matches[3]) == 0 {
			end = uint64(65535)
		} else {
			end, err = strconv.ParseUint(matches[3], 10, 32)
			if err != nil {
				fmt.Println("error: invalid port range specification")
				os.Exit(1)
			}
		}
	}

	return uint(begin), uint(end)
}

func Run() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println(usage)
		os.Exit(1)
	}

	begin, end := parseRange(portRangeSpec)

	status := map[bool]string{false: "Down", true: "Up"}
	for _, hostname := range flag.Args() {
		up := host.Scan(hostname)
		fmt.Printf("Host: %s\tStatus: %s\n", hostname, status[up])
		if !up {
			continue
		}
		for portn := begin; portn <= end; portn += 1 {
			if port.Scan(hostname, portn) {
				fmt.Printf("Host: %s\tPorts: %d/open/tcp/\n", hostname, portn)
			}
		}
	}
}
