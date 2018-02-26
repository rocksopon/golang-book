package main

import (
	"fmt"
)

type VendingMachine struct {
	insertedMoney int

}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	if coin == "T" {
		m.insertedMoney += 10
	} else if coin == "F" {
		m.insertedMoney += 5
	} else if coin == "TW" {
		m.insertedMoney += 2
	} else if coin == "O" {
		m.insertedMoney += 1
	}
}

func main() {
	vm := VendingMachine{}
	fmt.Println("Inserted Money:", vm.InsertedMoney())

	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
}