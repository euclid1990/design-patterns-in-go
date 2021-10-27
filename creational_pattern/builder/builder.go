package main

const (
	HouseCondo    string = "Condo"
	HouseAparment        = "Apartment"
	GarageBig            = "Big"
	GarageSmall          = "Small"
)

type iBuilder interface {
	SetCategory(string) iBuilder
	SetWindow(int) iBuilder
	SetFloor(int) iBuilder
	SetGarage(string) iBuilder
	Build() iBuilder
}

type HouseBuilder struct {
	category     string
	numOfWindows int
	numOfFloors  int
	garage       string
}

func (h *HouseBuilder) SetCategory(category string) iBuilder {
	h.category = category
	return h
}

func (h *HouseBuilder) SetWindow(window int) iBuilder {
	h.numOfWindows = window
	return h
}

func (h *HouseBuilder) SetFloor(floor int) iBuilder {
	if h.category == HouseAparment {
		panic("Can not set Floor to Apartment")
	}
	h.numOfFloors = floor
	return h
}

func (h *HouseBuilder) SetGarage(size string) iBuilder {
	if h.category == HouseAparment {
		panic("Can not set Garage to Apartment")
	}
	h.garage = size
	return h
}

func (h *HouseBuilder) Build() iBuilder {
	return h
}

var NewHouseBuilder = func() iBuilder {
	return &HouseBuilder{}
}
