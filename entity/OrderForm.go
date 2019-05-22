package entity

import "time"

type OrderForm struct {
	Id            uint      `json:"id"`
	BookName      string    `json:"book_name"`
	SellerId      uint      `json:"seller_id"`
	BuyerId       uint      `json:"buyer_id"`
	Price         float32   `json:"price"`
	CreateTime    time.Time `json:"create_time"`
	TradingMethod uint      `json:"trading_method"`
	Status        uint      `json:"status"`
}
