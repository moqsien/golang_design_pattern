package builder

// 建造者模式
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *ManufacturingDirector) Constructor() {
	m.builder.SetSeats().SetWheels().SetStructure()
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}
