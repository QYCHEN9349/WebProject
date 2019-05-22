package dao

import (
	"log"
	"webProject/constant"
	"webProject/entity"
	"webProject/webUtil"
)

type BookDao struct {
}

func (p *BookDao) Insert(book *entity.Book) int64 {
	result, err := webUtil.DB.Exec("INSERT INTO	book('book_name','seller_id','catagroy','original_price',"+
		"'selling_price','description','cover_page','link','status') "+"VALUES(?,?,?,?,?,?,?,?)",
		book.BookName, book.SellerId, book.Catagory, book.OriginalPrice, book.SellingPrice,
		book.Description, book.CoverPage, book.Link, book.Status)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}

func (p *BookDao) FindBookById(id uint) entity.Book {

	var book entity.Book
	row := webUtil.DB.QueryRow("SELECT * FROM book WHERE 'id' = ? ", id)
	err := row.Scan(&book.Id, &book.BookName, &book.SellerId, &book.Catagory, &book.OriginalPrice, &book.SellingPrice,
		&book.Description, &book.CoverPage, &book.Link, book.Status, book.CreateTime)
	if err != nil {
		log.Println(err)
	}
	return book
}

func (p *BookDao) FindBooksBySellerId(sellerId uint, pageNum uint) []entity.Book {
	rows, err := webUtil.DB.Query("SELECT * FROM book WHERE 'seller_id'=? LIMIT ? OFFSET ?",
		sellerId, constant.MaxBooksOfOnePage, constant.MaxBooksOfOnePage*pageNum)
	if err != nil {
		log.Println(err)
		return nil
	}
	var books []entity.Book
	for rows.Next() {
		var book entity.Book
		err := rows.Scan(&book.Id, &book.BookName, &book.SellerId, &book.Catagory, &book.OriginalPrice, &book.SellingPrice,
			&book.Description, &book.CoverPage, &book.Link, book.Status, book.CreateTime)
		if err != nil {
			log.Println(err)
			continue
		}
		books = append(books, book)
	}
	rows.Close()
	return books
}
func (p *BookDao) Update(book entity.Book) int64 {
	result, err := webUtil.DB.Exec("UPDATE book SET 'book_name'=?,'seller_id'=?,'catagroy'=?,"+
		"'original_price=?','selling_price=?','description=?','cover_page=?','link=?','status=?'",
		book.BookName, book.SellerId, book.Catagory, book.OriginalPrice, book.SellingPrice,
		book.Description, book.CoverPage, book.Link, book.Status)
	if err != nil {
		log.Println(err)
		return 0
	}
	num, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return num
}
func (p *BookDao) FindBooksByCondition(pageNum uint, bookName string, catagory string, maxPrice float32, minPrice float32) []entity.Book {
	sql := "SELECT * FROM book WHERE 1=1"
	if bookName != "" {
		sql = sql + "AND 'book_name'=" + bookName
	}
	if catagory != "" {
		sql = sql + "AND 'catagory'=" + catagory
	}
	sql = sql + "AND 'selling_price' < ? AND 'selling_price' > ? LIMIT ? OFFSET ?"
	rows, err := webUtil.DB.Query(sql, maxPrice, minPrice, constant.MaxBooksOfOnePage, constant.MaxBooksOfOnePage*pageNum)
	if err != nil {
		log.Println(err)
		return nil
	}
	var books []entity.Book
	for rows.Next() {
		var book entity.Book
		err := rows.Scan(&book.Id, &book.BookName, &book.SellerId, &book.Catagory, &book.OriginalPrice, &book.SellingPrice,
			&book.Description, &book.CoverPage, &book.Link, book.Status, book.CreateTime)
		if err != nil {
			log.Println(err)
			continue
		}
		books = append(books, book)
	}
	rows.Close()
	return books
}
func (p *BookDao) FindAllBooks(pageNum uint) []entity.Book {
	rows, err := webUtil.DB.Query("SELECT * FROM book LIMIT ? OFFSET ?",
		constant.MaxBooksOfOnePage, constant.MaxBooksOfOnePage*pageNum)
	if err != nil {
		log.Println(err)
		return nil
	}
	var books []entity.Book
	for rows.Next() {
		var book entity.Book
		err := rows.Scan(&book.Id, &book.BookName, &book.SellerId, &book.Catagory, &book.OriginalPrice, &book.SellingPrice,
			&book.Description, &book.CoverPage, &book.Link, book.Status, book.CreateTime)
		if err != nil {
			log.Println(err)
			continue
		}
		books = append(books, book)
	}
	rows.Close()
	return books
}
