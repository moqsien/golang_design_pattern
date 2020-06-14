package bridge

import (
	"errors"
	"fmt"
	"io"
)

// 桥接模式
type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}
type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(p.Writer, "%s", msg)
	return nil
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

type PacketPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (n *NormalPrinter) Print() error {
	n.Printer.PrintMessage(n.Msg)
	return nil
}

func (p *PacketPrinter) Print() error {
	p.Printer.PrintMessage(fmt.Sprintln("Packet Message: %s", p.Msg))
	return nil
}
