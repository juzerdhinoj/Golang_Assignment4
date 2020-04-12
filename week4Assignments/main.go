package main

import (
	"fmt"
	"time"
)

type Car struct {
	Name         string
	Color        string
	Insurance_ID string
	Gear         int
}

type Owner struct {
	Name            string
	Age             string
	Driving_License int
	Cars            []Car
}

type Insurance struct {
	ID       string
	Validity string
	coverage int
	desc     string
}

var stock []Car = []Car{
	{
		Name:         "Swift",
		Color:        "Blue",
		Insurance_ID: "001",
		Gear:         5,
	},
	{
		Name:         "City",
		Color:        "Green",
		Insurance_ID: "002",
		Gear:         5,
	}, {
		Name:         "Laura",
		Color:        "White",
		Insurance_ID: "003",
		Gear:         6,
	},
}

//var owner1Car Car = Car{{Name: "Swift", Color: "Green", Insurance_ID: "006", Gear: 5}}

//var owner1 Owner = Owner{Name: "Raul", Age: "39", Driving_License: 223, Cars: [owner1Car]}

func main() {
	go pinger()

	AddCar()
	UpdateCar()
	QueryStock("Fortuner")

}

func pinger() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println("ping")
	}
}

func QueryStock(name string) {
	for _, car := range stock {
		if car.Name == name {
			fmt.Println("Car: ", car)
			return
		}
	}
	fmt.Println("No such Car")
}

func UpdateCar() {
	var name, color, insuranceID string
	var gear int

	fmt.Println("Enter Car name you want to update:")
	fmt.Scanln(&name)
	fmt.Println("Enter new Car Color:")
	fmt.Scanln(&color)
	fmt.Println("Enter new Car Insurance ID:")
	fmt.Scanln(&insuranceID)
	fmt.Println("Enter new Car Gear:")
	fmt.Scanln(&gear)

	for _, car := range stock {
		if car.Name == name {

			car.Color = color
			car.Insurance_ID = insuranceID
			car.Gear = gear
			fmt.Println("New Car Details:  ", car)
			return
		}
	}
	fmt.Println("No such Car")

}

func AddCar() {

	var name, color, insuranceID string
	var gear int

	fmt.Println("Enter Car name:")
	fmt.Scanln(&name)
	fmt.Println("Enter Car Color:")
	fmt.Scanln(&color)
	fmt.Println("Enter Car Insurance ID:")
	fmt.Scanln(&insuranceID)
	fmt.Println("Enter Car Gear:")
	fmt.Scanln(&gear)

	var newCar Car = Car{Name: name, Color: color, Insurance_ID: insuranceID, Gear: gear}
	stock = append(stock, newCar)
}

//func to sell car to owner
func SellCar(owner Owner) {

	var name string

	fmt.Println("Enter Car name you want to sell:")
	fmt.Scanln(&name)

	for _, car := range stock {
		if car.Name == name {
			owner.Cars = append(owner.Cars, car)
			fmt.Println("Car Sold: ", car)
			return
		}
	}
	fmt.Println("No such Car")

}

//func to buy car from Owner
func BuyCar(owner Owner) {
	var name string

	fmt.Println("Enter Car name you want to Buy:")
	fmt.Scanln(&name)

	for _, car := range owner.Cars {
		if car.Name == name {
			stock = append(stock, car)
			fmt.Println("Car Bought: ", car)
			return
		}
	}
	fmt.Println("No such Car")
}
