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
	cons := &console.Console{Table: tabwriter.NewWriter(os.Stdout, 8, 2, 2, '\t', 0)}
	nData := new(console.Data)
	oData := new(console.Data)
	event := false
	firstRun := true
	s, err := vmguestlib.NewSession()

	// Probably missing the VMware tools
	if err != nil {
		exit(err)
	}

	fields := make(map[string]func(c *console.Console, n *console.Data, o *console.Data, s *vmguestlib.Session))

	fields["Time"] = console.PrintCurrentTime

	if conf.Guest {
		fields["CStlG"] = console.PrintCPUStolen
		fields["CUseG"] = console.PrintCPUUsed
	}
	if conf.Host {

	}
	if conf.Cpu {

	}
	if conf.Mem {

	}

	// Print table headers
	for field, _ := range fields {
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
			for _, call := range fields {
				call(cons, nData, oData, s)
			}
		} else {
			firstRun = false
		}

	End:
		// Increment refresh counter, save data then sleep
		if (conf.Count != 0) || (count == conf.Count-1) {
			break
		} else {
			count += 1
		}

		copier.Copy(oData, nData)
		time.Sleep(conf.Delay)

	}

}
