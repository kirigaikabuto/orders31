package orders31

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
)

type OrderAmqpEndpoints struct {
	ordersStore OrdersStore
}

func NewProductAmqpEndpoints(o OrdersStore) OrderAmqpEndpoints {
	return OrderAmqpEndpoints{ordersStore: o}
}

func (o *OrderAmqpEndpoints) CreateOrderAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		order := &Order{}
		jsonData := message.Body
		err := json.Unmarshal(jsonData, &order)
		if err != nil {
			panic(err)
		}
		newProduct, err := o.ordersStore.Create(order)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(newProduct)
		if err != nil {
			panic(err)
		}
		return &amqp.Message{Body: response}
	}
}

func (o *OrderAmqpEndpoints) ListOrderAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		order := &Order{}
		jsonData := message.Body
		err := json.Unmarshal(jsonData, &order)
		if err != nil {
			panic(err)
		}
		orders, err := o.ordersStore.ListOrdersByUserId(order.UserId)
		response, err := json.Marshal(orders)
		if err != nil {
			panic(err)
		}
		return &amqp.Message{Body: response}
	}
}
