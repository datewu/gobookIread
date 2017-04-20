package abstractFactory

// SportMotorbike implement Vehicle interface
type SportMotorbike struct{}

// GetMotorbikeType statisfy Motorbike intreface
func (*SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}

// NumWheels statisfy Vehicle interface
func (*SportMotorbike) NumWheels() int {
	return 2
}

// NumSeats statisfy Vehicle interface
func (*SportMotorbike) NumSeats() int {
	return 1
}
