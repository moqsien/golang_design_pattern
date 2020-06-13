package adapter

import (
	"fmt"
)

type LegacyPrinter interface {
	Print(s string) string
}

type ModernPrinter interface {
	PrintSorted() string
}

type LegacyPrinterImp struct{}

func (l *LegacyPrinterImp) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	fmt.Println(newMsg)
	return
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintSorted() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s\n", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
	return
}
