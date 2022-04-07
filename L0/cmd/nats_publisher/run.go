package main

import (
	"encoding/json"
	"fmt"
	"github.com/wubba-com/L0/internal/app/domain"
	"github.com/wubba-com/L0/internal/config"
	"github.com/wubba-com/L0/pkg/nats"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

const (
	clientID = "test"
)

func main() {
	cfg := config.GetConfig()
	order := &domain.Order{}

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("[err] nats-pub os: %s\n", err.Error())
		return
	}
	log.Println(wd)

	f, err := os.Open(filepath.Join(wd, cfg.DataFile))
	if err != nil {
		log.Printf("[err] nats-pub open file: %s\n", err.Error())
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("[err] nats-pub ioutil: %s\n", err.Error())
		return
	}
	err = json.Unmarshal(b, &order)
	if err != nil {
		log.Printf("[err] nats-pub JSON: %s\n", err.Error())
		return
	}

	sc := nats.NewStanConn(cfg.Nats.ClusterID, clientID) //

	// generate count orders
	orders := GenerateOrders(12, *order)
	fmt.Println("order: ", len(orders), orders[0].OrderUID, orders[0].CustomerID, order.Delivery.Name)
	for _, order := range orders {

		for _, item := range order.Items {
			uid := rand.Uint32()
			item.ChrtID = uint64(uid)
		}
		fmt.Println("ChID:", order.Items[0].ChrtID)
		b, err := json.Marshal(order)
		if err != nil {
			log.Printf("[err] nats-pub: %s\n", err.Error())
			return
		}

		//if i%2 == 0 {
		//	badJson := `{"ID":"test123", "fields":"lalala"}`
		//	err = sc.Publish(cfg.Nats.Channel, []byte(badJson))
		//	if err != nil {
		//		log.Printf("[err] nats-pub: %s\n", err.Error())
		//		return
		//	}
		//}

		err = sc.Publish(cfg.Nats.Channel, b)
		if err != nil {
			log.Printf("[err] nats-pub: %s\n", err.Error())
			return
		}
		//time.Sleep(1 * time.Second)
	}

	err = sc.Close()
	if err != nil {
		log.Printf("[err] nats-pub: %s\n", err.Error())
		return
	}
}

// GenerateOrders generates orders
func GenerateOrders(countOrders int, order domain.Order) []*domain.Order {
	orders := make([]*domain.Order, 0)

	for countOrders > 0 {
		o := order
		randInt := fmt.Sprintf("uid-%d", rand.Int())
		o.OrderUID = randInt

		orders = append(orders, &o)
		countOrders--
	}

	return orders
}
