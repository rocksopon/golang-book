package main

import (
	"fmt"
)

type NewVendingMachine struct {
	coins int
	items string
}

func main() {
	vm := NewVendingMachine{}

	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney())

	can := vm.SelectSD()
	fmt.Println(can)
}

func (n *NewVendingMachine) InsertCoin(coins string) {
	if coins == "T" {
		n.coins += 10
	} else if coins == "F" {
		n.coins += 5
	} else if coins == "TW" {
		n.coins += 2
	} else {
		n.coins += 1
	}
}

func (n *NewVendingMachine) GetInsertedMoney() int {
	return n.coins
}

func (n *NewVendingMachine) SelectSD() string {
	return "SD"
}

func (n *NewVendingMachine) SelectCC() string {
	return "CC"
}

func (n *NewVendingMachine) CoinReturn() string {
	return "SD"
}


