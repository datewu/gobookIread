package abstractFactory

// CruiseMotorbike implement Vehicle interface
type CruiseMotorbike struct{}

// GetMotorbikeType statisfy Motorbike intreface
func (*CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}

// NumWheels statisfy Vehicle interface
func (*CruiseMotorbike) NumWheels() int {
	return 2
}

// NumSeats statisfy Vehicle interface
func (*CruiseMotorbike) NumSeats() int {
	return 2
}
