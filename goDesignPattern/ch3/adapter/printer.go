package main

import "fmt"

type legacyPrinter interface {
	printa(s string) string
}

type myLegacyPrinter struct{}

func (myLegacyPrinter) printa(s string) string {
	return fmt.Sprintf("Legacy Printer: %s\n", s)
}

type modernPrinter interface {
	printStored() string
}

type printerAdapter struct {
	oldPrinter legacyPrinter
	msg        string
}

func (p *printerAdapter) printStored() string {
	if p.oldPrinter != nil {
		return p.oldPrinter.printa("Adapter: " + p.msg)
	}
	return p.msg
}
