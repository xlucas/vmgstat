package console

import (
	"fmt"
	"text/tabwriter"
	"time"

	"github.com/mgutz/ansi"
)

type Console struct {
	Table *tabwriter.Writer
	Color bool
}

func (c *Console) WriteHeaderCol(name string) {
	if c.Color {
		fmt.Fprintf(c.Table, "%s %7s %s\t", ansi.ColorCode("yellow+b"), name, ansi.ColorCode("reset"))
	} else {
		fmt.Fprintf(c.Table, "%s\t", name)
	}
}

func (c *Console) WriteLineEnd() {
	fmt.Fprintf(c.Table, "\n")
	c.Table.Flush()
}

func (c *Console) WritePercentCol(value float64) {
	fmt.Fprintf(c.Table, "%5.2f\t", value)
}

func (c *Console) WriteFloat64(value float64) {
	if value < 100000.0 {
		c.WritePercentCol(value)
	} else {
		fmt.Fprintf(c.Table, "%3.1e\t", value)
	}
}

func (c *Console) WriteUint32(value uint32) {
	if value < 1000000000 {
		fmt.Fprintf(c.Table, "%d\t", value)
	} else {
		fmt.Fprintf(c.Table, "%3.1e\t", float64(value))
	}
}

func (c *Console) WriteString(msg string) {
	fmt.Fprintf(c.Table, "%s\t", msg)
}

func (c *Console) WriteTimeCol(value time.Time) {
	if c.Color {
		fmt.Fprintf(c.Table, "%s%02d:%02d:%02d%s\t", ansi.ColorCode("yellow+b"), value.Hour(), value.Minute(), value.Second(), ansi.ColorCode("reset"))
	} else {
		fmt.Fprintf(c.Table, "%02d:%02d:%02d\t", value.Hour(), value.Minute(), value.Second())
	}
}

func (c *Console) WriteNaCol() {
	fmt.Fprintf(c.Table, " N/A \t")
}
