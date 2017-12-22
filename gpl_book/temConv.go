package main

type celsius float64

type fahrenheit float64

const (
	absoluteZeroC celsius = -273.15
	freezingC     celsius = 0
	boilingC      celsius = 100
)

func cTof(c celsius) fahrenheit {
	return fahrenheit(c*9/5 + 32)
}

func fToc(f fahrenheit) celsius {
	return celsius((f - 32) * 5 / 9)
}

func main() {

}
