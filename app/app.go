package main

import (
	"ck/app/orderer"
)

func main() {
	orderer := orderer.New("./app/orders.json")
	orderer.PlaceOrder()
}
