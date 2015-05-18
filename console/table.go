package console

import (
	"fmt"
	"text/tabwriter"
	"time"
)

type Console struct {
	Table *tabwriter.Writer
}

func (c *Console) WriteHeaderCol(name string) {
	fmt.Fprintf(c.Table, "%s\t", name)
}

func (c *Console) WriteLineEnd() {
	fmt.Fprintf(c.Table, "\n")
	c.Table.Flush()
}

func (c *Console) WritePercentCol(value float64) {
	fmt.Fprintf(c.Table, "%5.2f\t", value)
}

func (c *Console) WriteFloat64(value float64) {
	if value < 1000000.0 {
		c.WritePercentCol(value)
	} else {

		fmt.Fprintf(c.Table, "%3.1e\t", value)
	}
}

func (c *Console) WriteUint32(value uint32) {
	if value < 1000000.0 {
		c.WritePercentCol(float64(value))
	} else {
		fmt.Fprintf(c.Table, "%3.1e\t", float64(value))
	}
}

func (c *Console) WriteTimeCol(value time.Time) {
	fmt.Fprintf(c.Table, "%02d:%02d:%02d\t", value.Hour(), value.Minute(), value.Second())
}

func (c *Console) WriteNaCol() {
	fmt.Fprintf(c.Table, " N/A \t")
}
