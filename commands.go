package orders

type CreateOrderCommand struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"-"`
}

type HttpError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
