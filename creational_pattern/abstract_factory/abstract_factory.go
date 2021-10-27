package main

import (
	"fmt"
	"log"
	"strings"
)

type OsType int

const (
	WINDOWS OsType = iota
	MAC
)

type IGUIFactory interface {
	createButton() IButton
	createCheckbox() ICheckbox
}

type IButton interface {
	paint() string
}

type ICheckbox interface {
	paint() string
}

type WinFactory struct {
	os OsType
}

func (wF *WinFactory) createButton() IButton {
	return new(WinButton)
}

func (wF *WinFactory) createCheckbox() ICheckbox {
	return new(WinCheckbox)
}

type WinButton struct{}

func (wB *WinButton) paint() string {
	log.Println("Windows_Button")
	return "Windows_Button"
}

type WinCheckbox struct{}

func (wB *WinCheckbox) paint() string {
	log.Println("Windows_Checkbox")
	return "Windows_Checkbox"
}

type MacFactory struct {
	os OsType
}

func (mF *MacFactory) createButton() IButton {
	return new(MacButton)
}

func (mF *MacFactory) createCheckbox() ICheckbox {
	return new(MacCheckbox)
}

type MacButton struct{}

func (wB *MacButton) paint() string {
	log.Println("Mac_Button")
	return "Mac_Button"
}

type MacCheckbox struct{}

func (wB *MacCheckbox) paint() string {
	log.Println("Mac_Checkbox")
	return "Mac_Checkbox"
}

type Application struct {
	factory IGUIFactory
}

func (a *Application) SetOsFactory(os OsType) (*Application, error) {
	switch os {
	case WINDOWS:
		a.factory = new(WinFactory)
	case MAC:
		a.factory = new(MacFactory)
	default:
		return nil, fmt.Errorf("Wrong OS type passed")
	}
	return a, nil
}

func (a *Application) SetFactory(f IGUIFactory) {
	a.factory = f
}

func (a *Application) GetFactory() IGUIFactory {
	return a.factory
}

func (a *Application) CreateUI() string {
	button := a.factory.createButton()
	checkbox := a.factory.createCheckbox()
	log.Println("Render UI")
	b, c := button.paint(), checkbox.paint()
	return strings.Join([]string{b, c}, "-")
}
