package orders31

type OrdersStore interface {
	Create(order *Order) (*Order, error)
	ListOrdersByUserId(userId string) ([]Order, error)
}
