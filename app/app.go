package main

import (
	"github.com/srnewbie/ck/app/orderer"
)

func main() {
	orderer := orderer.New("./app/orders.json")
	orderer.PlaceOrder()
}
