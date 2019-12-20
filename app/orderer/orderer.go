package orderer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"ck/models"
)

type Orderer struct {
	Orders models.Orders
	timer  *time.Timer
}

func New(p string) *Orderer {
	f, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bytes, _ := ioutil.ReadAll(f)
	orderer := &Orderer{timer: time.NewTimer(0)}
	json.Unmarshal(bytes, &orderer.Orders)
	return orderer
}

func (o *Orderer) PlaceOrder() {
	rand.Seed(time.Now().UnixNano())
	o.Delay()
	<-o.timer.C
	for _, order := range o.Orders {
		<-o.timer.C
		fmt.Println("placing order:")
		fmt.Println(order.ID, ",", order.Name, ",", order.Type, ",", order.OrderTime)
		o.Post(order)
		o.Delay()
	}
}

func (o *Orderer) Post(order *models.Order) {
	v, _ := json.Marshal(order)
	resp, err := http.Post(
		"http://127.0.0.1:3000",
		"application/json",
		bytes.NewBuffer(v),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func (o *Orderer) Delay() {
	t := rand.Intn(10)
	fmt.Println("next order coming in", t, "seconds")
	o.timer.Reset(time.Duration(t) * time.Second)
}
