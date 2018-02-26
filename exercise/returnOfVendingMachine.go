package main

import (
	"fmt"
)

type VendingMachine struct {
	insertedMoney int
	coins map[string]int

}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney += m.coins[coin]
	/*if coin == "T" {
		m.insertedMoney += 10
	} else if coin == "F" {
		m.insertedMoney += 5
	} else if coin == "TW" {
		m.insertedMoney += 2
	} else if coin == "O" {
		m.insertedMoney += 1
	}*/
}

func main() {
	var coins = map[string]int{"T": 1, "F": 5, "TW":2, "O":1}
	vm := VendingMachine{coins: coins}
	fmt.Println("Inserted Money:", vm.InsertedMoney())

	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
}