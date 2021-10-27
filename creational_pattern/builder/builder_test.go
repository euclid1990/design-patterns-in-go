package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) SetupTest() {
}

func (suite *TestSuite) TestHouseBuilderSetCategory() {
	houseBuilder := NewHouseBuilder()
	ptr := reflect.ValueOf(houseBuilder.SetCategory(HouseCondo))
	v := reflect.Indirect(ptr)
	if v.Kind() == reflect.Struct {
		category := v.FieldByName("category")
		assert.Equal(suite.T(), category.String(), HouseCondo)
	}
}

func (suite *TestSuite) TestHouseBuilderSetWindow() {
	houseBuilder := NewHouseBuilder()
	ptr := reflect.ValueOf(houseBuilder.SetWindow(3))
	v := reflect.Indirect(ptr)
	if v.Kind() == reflect.Struct {
		numOfWindows := v.FieldByName("numOfWindows")
		assert.Equal(suite.T(), int(numOfWindows.Int()), 3)
	}
}

func (suite *TestSuite) TestHouseBuilderSetFloor() {
	houseBuilder := NewHouseBuilder()
	ptr := reflect.ValueOf(houseBuilder.SetFloor(3))
	v := reflect.Indirect(ptr)
	if v.Kind() == reflect.Struct {
		numOfFloors := v.FieldByName("numOfFloors")
		assert.Equal(suite.T(), int(numOfFloors.Int()), 3)
	}
}

func (suite *TestSuite) TestHouseBuilderSetGarage() {
	houseBuilder := NewHouseBuilder()
	ptr := reflect.ValueOf(houseBuilder.SetGarage(GarageBig))
	v := reflect.Indirect(ptr)
	if v.Kind() == reflect.Struct {
		garage := v.FieldByName("garage")
		assert.Equal(suite.T(), garage.String(), GarageBig)
	}
}

func (suite *TestSuite) TestHouseBuilderBuild() {
	OriginNewHouseBuilder := NewHouseBuilder
	NewHouseBuilder = func() iBuilder {
		return &HouseBuilder{}
	}

	expect := &HouseBuilder{
		category:     HouseCondo,
		numOfWindows: 5,
		numOfFloors:  6,
		garage:       GarageSmall,
	}
	condoHouse := NewHouseBuilder().
		SetCategory(HouseCondo).
		SetWindow(5).
		SetFloor(6).
		SetGarage(GarageSmall).
		Build()
	assert.Equal(suite.T(), condoHouse, expect)
	NewHouseBuilder = OriginNewHouseBuilder
}

func (suite *TestSuite) TestHouseBuilderBuildPanic() {
	expect := "Can not set Floor to Apartment"

	assert.Panics(suite.T(), func() {
		NewHouseBuilder().
			SetCategory(HouseAparment).
			SetWindow(5).
			SetFloor(6).
			Build()
	}, expect)

	expect = "Can not set Garage to Apartment"
	assert.Panics(suite.T(), func() {
		NewHouseBuilder().
			SetCategory(HouseAparment).
			SetWindow(5).
			SetGarage(GarageSmall).
			Build()
	}, expect)
}

func (suite *TestSuite) TestNewHouseBuilder() {
	assert.Equal(suite.T(), NewHouseBuilder(), &HouseBuilder{})
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
