package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/xlucas/go-vmguestlib/vmguestlib"
)

func main() {
	var err error
	var event bool
	var firstRun bool = true
	var s *vmguestlib.Session

	var guestFlag bool
	var hostFlag bool
	var cpuFlag bool
	var memFlag bool
	var count uint64
	var delay time.Duration
	var currentCount uint64 = 0

	flag.BoolVar(&guestFlag, "guest", true, "Fetch Guest statistics.")
	flag.BoolVar(&hostFlag, "host", false, "Fetch Host statistics.")
	flag.BoolVar(&cpuFlag, "cpu", true, "Fetch CPU statistics.")
	flag.BoolVar(&memFlag, "memory", true, "Fetch Memory statistics.")
	flag.Uint64Var(&count, "count", 5, "Refresh count")
	flag.DurationVar(&delay, "delay", 1*time.Second, "Refresh delay.")
	flag.Parse()

	if s, err = vmguestlib.NewSession(); err != nil {
		fmt.Fprintln(os.Stderr, "An error occured while opening a new session : %s", err)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 2, 2, ' ', 0)

	var nStealG, nUsedG, nElapsed uint64
	var oStealG, oUsedG, oElapsed uint64

	for {

		if currentCount == count {
			break
		}

		if event, err = s.RefreshInfo(); err != nil {
			fmt.Fprintln(os.Stderr, "An error occured while refreshing statistics : %s", err)
		}

		if !event && !firstRun {
			if guestFlag {
				fmt.Fprintln(w, "Date\tStealG\tUsedG\t")

				if nStealG, err = s.GetCPUStolen(); err != nil {
					fmt.Println(os.Stderr, err)
				}
				if nUsedG, err = s.GetCPUUsed(); err != nil {
					fmt.Println(os.Stderr, err)
				}
				if nElapsed, err = s.GetTimeElapsed(); err != nil {
					fmt.Println(os.Stderr, err)
				}

				//fmt.Fprintln(os.Stdout, "Steal: ", stealG)
				//fmt.Fprintln(os.Stdout, "Used : ", usedG)
				//fmt.Fprintln(os.Stdout, "Elaspsed : ", elapsed)

				fmt.Fprintf(w, "%02d:%02d:%02d\t%3.1f\t%3.1f\t",
					time.Now().Hour(), time.Now().Minute(), time.Now().Second(),
					float64((nStealG-oStealG)/(nElapsed-oElapsed)*100.0),
					float64((nUsedG-oUsedG)/(nElapsed-oElapsed)*100.0),
				)

				oStealG = nStealG
				oUsedG = nUsedG
				oElapsed = nElapsed
			}
		}

		// Sleep and update counter
		time.Sleep(delay)
		currentCount += 1
		if firstRun {
			firstRun = false
		}

	}

}
