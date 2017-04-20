package abstractFactory

// LuxuryCar implement Vehicle interface
type LuxuryCar struct{}

// NumDoors statisfy Car interface
func (*LuxuryCar) NumDoors() int {
	return 4
}

// NumWheels statisfy Vehicle interface
func (*LuxuryCar) NumWheels() int {
	return 4
}

// NumSeats statisfy Vehicle interface
func (*LuxuryCar) NumSeats() int {
	return 5
}
