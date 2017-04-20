package abstractFactory

// Vehicle all objects in factories must implement
type Vehicle interface {
	NumWheels() int
	NumSeats() int
}
