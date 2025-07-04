package model

type Comment struct {
	ID        int    `json:"id"`
	ProductId int    `json:"product_id"`
	UserId    int    `json:"user_id"`
	Message   string `json:"message"`
}
