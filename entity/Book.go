package entity

import "time"

type Book struct {
	Id            uint    `json:"id"`
	BookName      string  `json:"book_name"`
	SellerId      uint    `json:"sell_id"`
	Catagory      string  `json:"catagory"`
	OriginalPrice float32 `json:"original_price"`
	SellingPrice  float32 `json:"selling_price"`
	Description   string  `json:"description"`
	CoverPage     []byte  `json:"cover_page"`
	Link          string  `json:"link"`
	//0-delete,1-onsale,2-sold
	Status     uint      `json:"status"`
	CreateTime time.Time `json:"create_time"`
}
