package datafeeder

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
	"gonum.org/v1/gonum/stat/distuv"
)

func RandomDataProducer() {
	dist := distuv.Normal{Mu: 1.0, Sigma: 1.0}
	nc, _ := nats.Connect(nats.DefaultURL)

	for {
		price := dist.Rand()
		pricestr := fmt.Sprintf("%v", price)
		nc.Publish("prices", []byte(pricestr))
	}
}
