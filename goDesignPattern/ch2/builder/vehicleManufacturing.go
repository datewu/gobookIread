package creational

// BuildProcess interface
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}

// ManufacturingDirector in charge of construction of the objects
type ManufacturingDirector struct {
	builder BuildProcess
}

// SetBuilder set the builders being used in the ManufacturingDirector
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Construct use the builders and reproduce the required steps
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// VehicleProduct the result
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// CarBuilder build cars
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels statisfy BuildProcess interface
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetSeats statisfy BuildProcess interface
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetStructure statisfy BuildProcess interface
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

// Build statisfy BuildProcess interface
func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}

// BikeBuilder build cars
type BikeBuilder struct {
	v VehicleProduct
}

// SetWheels statisfy BuildProcess interface
func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}

// SetSeats statisfy BuildProcess interface
func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 2
	return c
}

// SetStructure statisfy BuildProcess interface
func (c *BikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Motorbike"
	return c
}

// Build statisfy BuildProcess interface
func (c *BikeBuilder) Build() VehicleProduct {
	return c.v
}
