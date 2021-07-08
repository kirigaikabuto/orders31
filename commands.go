package orders31

type CreateOrderCommand struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}

type HttpError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
