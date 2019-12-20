package orderer

import (
	"bytes"
	"ck/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	orderer := &Orderer{timer: time.NewTimer(1)}
	json.Unmarshal(bytes, &orderer.Orders)
	return orderer
}

func (o *Orderer) PlaceOrder() {
	for i := 0; i < len(o.Orders); i += 3 {
		<-o.timer.C
		fmt.Println("placing order:")
		j := i + 3
		if j > len(o.Orders) {
			j = len(o.Orders)
		}
		o.Post(o.Orders[i:j])
		o.timer.Reset(time.Duration(1) * time.Second)
	}
}

func (o *Orderer) Post(orders []*models.Order) {
	for _, o := range orders {
		util.DisplayOrder(o)
	}

	v, _ := json.Marshal(orders)
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
