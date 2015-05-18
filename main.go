package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/jinzhu/copier"
	"github.com/xlucas/go-vmguestlib/vmguestlib"
	"github.com/xlucas/vmgstat/cli"
	"github.com/xlucas/vmgstat/console"
)

func exit(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "An error occured : ", err)
	}
	os.Exit(1)
}

func main() {
	conf := cli.Config{}

	flag.BoolVar(&conf.Guest, "guest", true, "Show guest information.")
	flag.BoolVar(&conf.Host, "host", false, "Show host information.")
	flag.BoolVar(&conf.Cpu, "cpu", false, "Show cpu stats.")
	flag.BoolVar(&conf.Mem, "mem", false, "Show memory stats.")
	flag.UintVar(&conf.Count, "count", 0, "Refresh count.")
	flag.DurationVar(&conf.Delay, "delay", 1*time.Second, "Refresh delay.")
	flag.Parse()

	count := uint(0)
	cons := &console.Console{Table: tabwriter.NewWriter(os.Stdout, 10, 2, 0, ' ', 0)}
	nData, oData := new(console.Data), new(console.Data)
	event, firstRun := false, true
	fields := make(map[string]console.PrintFunc)
	order := make([]string, 0)

	// Open vSphere session
	s, err := vmguestlib.NewSession()
	// Probably missing the VMware tools
	if err != nil {
		exit(err)
	}

	// Append fields
	console.AppendField(&fields, &order, "Time", console.PrintCurrentTime)

	if conf.Guest {
		if conf.Cpu {
			console.AppendField(&fields, &order, "CLimG", console.PrintCPULimit)
			console.AppendField(&fields, &order, "CResG", console.PrintCPUReservation)
			console.AppendField(&fields, &order, "CShaG", console.PrintCPUShares)
			console.AppendField(&fields, &order, "CStlG", console.PrintCPUStolen)
			console.AppendField(&fields, &order, "CUseG", console.PrintCPUUsed)
		}
		if conf.Mem {

		}
	}
	if conf.Host {
		if conf.Cpu {
			console.AppendField(&fields, &order, "CUseH", console.PrintHostCPUUsed)
			console.AppendField(&fields, &order, "CNumH", console.PrintHostNumCPUCores)
			console.AppendField(&fields, &order, "CSpeH", console.PrintHostProcessorSpeed)
		}
		if conf.Mem {

		}
	}

	// Print table headers
	for _, field := range order {
		cons.WriteHeaderCol(field)
	}
	cons.WriteLineEnd()

	// Refresh until we reach the refresh count
	for {

		// Refresh session info
		if event, err = s.RefreshInfo(); err != nil {
			exit(err)
		}
		// Retrieve new data
		nData.Refresh(s)
		// vSphere event occured
		if event {
			fmt.Fprintln(os.Stdout, "--- VSPHERE GUEST SESSION CHANGED! WAIT FOR NEXT REFRESH ... ---")
			goto End
		}
		// Display field values
		if !firstRun {
			for _, field := range order {
				fields[field](cons, nData, oData, s)
			}
			cons.WriteLineEnd()
		} else {
			firstRun = false
		}

	End:
		// Increment refresh counter if needed
		if (conf.Count != 0) || (count == conf.Count-1) {
			break
		} else {
			count++
		}
		// Save data and sleep
		copier.Copy(oData, nData)
		time.Sleep(conf.Delay)

	}

}
