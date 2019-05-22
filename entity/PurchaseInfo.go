package entity

import "time"

type PurchaseInfo struct {
	Id          uint      `json:"id"`
	BookName    string    `json:"book_name"`
	BuyerId     uint      `json:"buyer_id"`
	Price       float32   `json:"price"`
	Description string    `json:"description"`
	Status      uint      `json:"status"`
	CreateTime  time.Time `json:"create_time"`
}
