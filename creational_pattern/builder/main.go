package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("WOHA! Program is panicking with msg:", r)
		}
	}()

	condo := NewHouseBuilder().
		SetCategory(HouseCondo).
		SetWindow(4).
		SetFloor(2).
		SetGarage(GarageBig).
		Build()
	fmt.Println("Condo House:", condo)

	apartment := NewHouseBuilder().
		SetCategory(HouseAparment).
		SetWindow(3).
		Build()
	fmt.Println("Apartment House:", apartment)

	apartmentErr := NewHouseBuilder().
		SetCategory(HouseAparment).
		SetWindow(3).
		SetFloor(2).
		Build()
	fmt.Println("Apartment House:", apartmentErr)
}
