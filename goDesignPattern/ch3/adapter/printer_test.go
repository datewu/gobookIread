package main

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World!"
	adapter := printerAdapter{&myLegacyPrinter{}, msg}
	rt := adapter.printStored()
	if rt != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didnot match: %s\n", rt)
	}

	adapter = printerAdapter{msg: msg}
	rt = adapter.printStored()
	if rt != "Hello World!" {
		t.Errorf("Message didnot match: %s\n", rt)
	}

}
