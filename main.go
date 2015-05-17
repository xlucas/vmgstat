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

	for {

		if currentCount == count {
			break
		}

		if event, err = s.RefreshInfo(); err != nil {
			fmt.Fprintln(os.Stderr, "An error occured while refreshing statistics : %s", err)
		}

		if !event {
			if guestFlag {
				var stealG, usedG, elapsed uint64
				//fmt.Fprintln(w, "Date\tStealG\tUsedG\t")

				if stealG, err = s.GetCPUStolen(); err != nil {
					fmt.Println(os.Stderr, err)
				}
				if usedG, err = s.GetCPUUsed(); err != nil {
					fmt.Println(os.Stderr, err)
				}
				if elapsed, err = s.GetTimeElapsed(); err != nil {
					fmt.Println(os.Stderr, err)
				}

				fmt.Println(os.Stdout, "Steal: ", stealG)
				fmt.Println(os.Stdout, "Used : ", usedG)
				fmt.Println(os.Stdout, "Elaspsed : ", elapsed)

				/*fmt.Fprintf(w, "%02d:%02d:%02d\t%3.1f\t%3.1f\t",
					time.Now().Hour(), time.Now().Minute(), time.Now().Second(),
					(stealG/elapsed)*100.0,
					(usedG/elapsed)*100.0,
				)*/

				w.Flush()
			}
		} else {
			// A vSphere event occured, statistics should be discarded until the next RefreshInfo() call
		}

		// Sleep and update counter
		time.Sleep(delay)
		currentCount += 1

	}

}
