package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockFactory struct {
	mock.Mock
}

func (m *MockFactory) createButton() IButton {
	args := m.Called()
	return args.Get(0).(IButton)
}

func (m *MockFactory) createCheckbox() ICheckbox {
	args := m.Called()
	return args.Get(0).(ICheckbox)
}

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) SetupTest() {
}

func (suite *TestSuite) TestWinFactoryCreateButton() {
	winFactory := new(WinFactory)
	button := winFactory.createButton()
	assert.Equal(suite.T(), "*main.WinButton", reflect.TypeOf(button).String())
}

func (suite *TestSuite) TestWinFactoryCreateCheckbox() {
	winFactory := new(WinFactory)
	checkbox := winFactory.createCheckbox()
	assert.Equal(suite.T(), "*main.WinCheckbox", reflect.TypeOf(checkbox).String())
}

func (suite *TestSuite) TestWinButtonPaint() {
	winButton := new(WinButton)
	str := winButton.paint()
	assert.Equal(suite.T(), "Windows_Button", str)
}

func (suite *TestSuite) TestWinCheckboxPaint() {
	winCheckbox := new(WinCheckbox)
	str := winCheckbox.paint()
	assert.Equal(suite.T(), "Windows_Checkbox", str)
}

func (suite *TestSuite) TestMacFactoryCreateButton() {
	macFactory := new(MacFactory)
	button := macFactory.createButton()
	assert.Equal(suite.T(), "*main.MacButton", reflect.TypeOf(button).String())
}

func (suite *TestSuite) TestMacFactoryCreateCheckbox() {
	macFactory := new(MacFactory)
	checkbox := macFactory.createCheckbox()
	assert.Equal(suite.T(), "*main.MacCheckbox", reflect.TypeOf(checkbox).String())
}

func (suite *TestSuite) TestMacButtonPaint() {
	macButton := new(MacButton)
	str := macButton.paint()
	assert.Equal(suite.T(), "Mac_Button", str)
}

func (suite *TestSuite) TestMacCheckboxPaint() {
	macCheckbox := new(MacCheckbox)
	str := macCheckbox.paint()
	assert.Equal(suite.T(), "Mac_Checkbox", str)
}

func (suite *TestSuite) TestApplicationSetOsFactory() {
	winFactory, _ := new(Application).SetOsFactory(WINDOWS)
	assert.Equal(suite.T(), "*main.WinFactory", reflect.TypeOf(winFactory.GetFactory()).String())
	macFactory, _ := new(Application).SetOsFactory(MAC)
	assert.Equal(suite.T(), "*main.MacFactory", reflect.TypeOf(macFactory.GetFactory()).String())
	_, err := new(Application).SetOsFactory(OsType(99))
	assert.Equal(suite.T(), err.Error(), "Wrong OS type passed")
}

func (suite *TestSuite) TestApplicationCreateUI() {
	mockFactory := new(MockFactory)
	app := new(Application)
	app.SetFactory(mockFactory)
	mockFactory.On("createButton").Return(&WinButton{})
	mockFactory.On("createCheckbox").Return(&WinCheckbox{})
	render := app.CreateUI()
	assert.Equal(suite.T(), render, "Windows_Button-Windows_Checkbox")
	mockFactory.AssertNumberOfCalls(suite.T(), "createButton", 1)
	mockFactory.AssertNumberOfCalls(suite.T(), "createCheckbox", 1)
	mockFactory.AssertExpectations(suite.T())
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
