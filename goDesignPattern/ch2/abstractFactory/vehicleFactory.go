package abstractFactory

import "fmt"

// VehicleFactory is the Abstract Factory
type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

// LuxuryCarType types
const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

// CarFactory implements the VehicleFactory interface
type CarFactory struct{}

// NewVehicle statisfy VehicleFactory interface
func (*CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}

// SportMotorbikeType types
const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

// MotorbikeFactory implements the VehicleFactory interface
type MotorbikeFactory struct{}

// NewVehicle statisfy VehicleFactory interface
func (*MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}

// CarFactoryType types
const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

// BuildFactory as the name suggest
func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)

	}
}
