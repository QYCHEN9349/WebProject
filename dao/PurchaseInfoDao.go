package dao

import (
	"log"
	"webProject/constant"
	"webProject/entity"
	"webProject/webUtil"
)

type PurchaseDaoInfo struct {
}

func (p *PurchaseDaoInfo) Insert(purchaseInfo *entity.PurchaseInfo) uint {
	result, err := webUtil.DB.Exec("INSERT INTO purchase_info('book_name','buyer_id','price','description','status')"+
		" VALUES(?.?.?,?,?)", purchaseInfo.BookName, purchaseInfo.BuyerId, purchaseInfo.Price, purchaseInfo.Description, purchaseInfo.Status)
	if err!=nil{
		log.Println(err)
	}
}

func (p *PurchaseDaoInfo) FindPurchaseInfo(pageNum uint) []entity.PurchaseInfo {
	rows, err := webUtil.DB.Query("SELECT * FROM purchase_info LIMIT ? OFFSET ?",
		constant.MaxBooksOfOnePage, pageNum*constant.MaxBooksOfOnePage)
	if err != nil {
		log.Println(err)
		return nil
	}
	var purchaseInfos []entity.PurchaseInfo
	for rows.Next() {
		var purchaseInfo entity.PurchaseInfo
		err := rows.Scan(&purchaseInfo.Id, &purchaseInfo.BookName, &purchaseInfo.BuyerId,
			&purchaseInfo.Price, &purchaseInfo.Description, &purchaseInfo.Status, &purchaseInfo.CreateTime)
		if err != nil {
			log.Println(err)
			continue
		}
		purchaseInfos = append(purchaseInfos, purchaseInfo)
	}
	return purchaseInfos
}

func (p *PurchaseDaoInfo) FindPurchaseInfoById(id uint, pageNum uint) []entity.PurchaseInfo {
	rows, err := webUtil.DB.Query("SELECT * FROM purchase_info WHERE 'id'=? LIMIT ? OFFSET ?",
		id, constant.MaxBooksOfOnePage, pageNum*constant.MaxBooksOfOnePage)
	if err != nil {
		log.Println(err)
		return nil
	}
	var purchaseInfos []entity.PurchaseInfo
	for rows.Next() {
		var purchaseInfo entity.PurchaseInfo
		err := rows.Scan(&purchaseInfo.Id, &purchaseInfo.BookName, &purchaseInfo.BuyerId,
			&purchaseInfo.Price, &purchaseInfo.Description, &purchaseInfo.Status, &purchaseInfo.CreateTime)
		if err != nil {
			log.Println(err)
			continue
		}
		purchaseInfos = append(purchaseInfos, purchaseInfo)
	}
	return purchaseInfos
}
func (p *PurchaseDaoInfo) update(purchaseInfo entity.PurchaseInfo) int64 {
	result, err := webUtil.DB.Exec("UPDATE purchase_info SET 'book_name'=?,'buyer_id'=?,'price'=?,'description'=?,'status'=?",
		purchaseInfo.BookName, purchaseInfo.BuyerId, purchaseInfo.Price, purchaseInfo.Description, purchaseInfo.Status)
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
