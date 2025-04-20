package printer

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	Out io.Writer
	Err io.Writer
}

func NewPrinter() *Printer {
	return &Printer{
		Out: os.Stdout,
		Err: os.Stderr,
	}
}

func (p *Printer) Info(msg string) {
	fmt.Fprintf(p.Out, "INFO: %s\n", msg)
}

func (p *Printer) Warn(msg string) {
	fmt.Fprintf(p.Err, "WARN: %s\n", msg)
}

func (p *Printer) Error(msg string) {
	fmt.Fprintf(p.Err, "ERR: %s\n", msg)
}
