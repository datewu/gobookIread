package abstractFactory

// FamilyCar implement Vehicle interface
type FamilyCar struct{}

// NumDoors statisfy Car interface
func (*FamilyCar) NumDoors() int {
	return 5
}

// NumWheels statisfy Vehicle interface
func (*FamilyCar) NumWheels() int {
	return 4
}

// NumSeats statisfy Vehicle interface
func (*FamilyCar) NumSeats() int {
	return 5
}
