package orders

type Order struct {
	Id        string `json:"id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}
