package webUtil

import (
	"database/sql"
	"log"
	"strings"
)

const (
	username   = "username"
	password   = "password"
	ip         = "127.0.0.1"
	port       = "3306"
	dbName     = "web"
	driverName = "mysql"
)

var DB *sql.DB

func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=uft8"
	//注意：要想解析time.Time类型，必须要设置parseTime=True
	path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", DBName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	//打开数据库，前者是驱动名，所以要导入:_"github.com/go-sql-driver/mysql"
	DB, _ = sql.Open(driverName, path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Panic(err)
	}
	log.Println("database connect success")
}

func CreateTable() {
	userTable := `CREATE TABLE IF NOT EXISTS 'user'(
				'id' INT UNSIGNED ATUO_INCREMENT,
				'username' VARCHAR(20) NOT NULL,
				'password' VARCHAR(20) NOT NULL,
				'phone_number' VARCHAR(20),
				'address' VARCHAR(100),
				PRIMARY KEY('id')
				);`
	BookTable := `CREATE TABLE IF NOT EXISTS 'book'(
				'id' INT UNSIGNED ATUO_INCREMENT,
				'bookname' VARCHAR(20) NOT NULL,
				'seller_id' INT UNSIGNED AUTO_INCREMENT,
				'catagory' VARCHAR(20) NOT NULL,
				'original_price' FLOAT(8,2),
				'selling_price' FLOAT(8,2) NOT NULL,
				'description' VARCHAR(200) NOT NULL,
				'cover_page' LONGBLOB,
				'link' VARCHAR(50),
				'status' ENUM('delete','onSale','sold'),
				'create_time' DATETIME,
				PRIMARY KEY('id')
				);`
	purchaseTable := `CREATE TABLE IF NOT EXISTS 'purchase_info'(
				'id' INT UNSIGNED ATUO_INCREMENT,
				'bookname' VARCHAR(20) NOT NULL,
				'buyer_id' INT UNSIHNED AUTO_INCREMENT,
				'price' FLOAT(8,2) NOT NULL,
				'description' VARCHAR(200) NOT NULL,
				'status' ENUM('delete','continue','finished'),
				'create_time' DATETIME,
				PRIMARY KEY('id')
				);`
	orderFormTable := `CREATE TABLE IF NOT EXISTS 'order_form'(
				'id' INT UNSIGNED ATUO_INCREMENT,
				'book_name' VARCHAR(20) NOT NULL,
				'seller_id' INT NOT NULL,
				'buyer_id' INT NOT NULL,
				'price' FLOAT(8,2) NOT NULL,
				'create_time' DATETIME,
				'trading_method' ENUM('online','underline')
				'status' ENUM('cancel','continue','finished'),
				PRIMARY KEY('id'),
				FOREIGN KEY('seller_id') REFERENCES user('id'),
				FOREIGN KEY('buyer_id') REFERENCES user('id')
				);`
	messageTable := `CREATE TABLE IF NOT EXISTS 'message'(
				'id' INT UNSIGNED ATUO_INCREMENT,
				'sender_id' INT NOT NULL,
				'receiver_id' INT NOT NULL,
				'content' VARCHAR(200) NOT NULL,
				'create_time' DATETIME,
				PRIMARY KEY('id'),
				FOREIGN KEY('sender_id') REFERENCES user('id'),
				FOREIGN KEY('receiver_id') REFERENCES user('id')
				);`

	_, err := DB.Query(userTable)
	if err != nil {
		log.Panic(err)
	}
	_, err = DB.Query(BookTable)
	if err != nil {
		log.Panic(err)
	}
	_, err = DB.Query(purchaseTable)
	if err != nil {
		log.Panic(err)
	}
	_, err = DB.Query(orderFormTable)
	if err != nil {
		log.Panic(err)
	}
	_, err = DB.Query(messageTable)
	if err != nil {
		log.Panic(err)
	}
}
