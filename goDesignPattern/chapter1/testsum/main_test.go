package main

import "testing"

func TestSum(t *testing.T) {
	a, b, expected := 5, 6, 11
	res := sum(a, b)
	if res != expected {
		t.Errorf("sum function doesnot work, %d+%d isnot %d\n", a, b, res)
	}

}

func TestMultiply(t *testing.T) {
	a, b, expected := 5, 7, 35

	res := multiply(a, b)
	if res != expected {
		t.Errorf("multiply function doesnot work, %d*%d isnot %d\n", a, b, res)
	}

}
