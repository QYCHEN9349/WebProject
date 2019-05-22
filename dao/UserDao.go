package dao

import (
	"log"
	"webProject/entity"
	"webProject/webUtil"
)

type UserDao struct {
}

func (p *UserDao) Insert(user *entity.User) int64 {
	result, err := webUtil.DB.Exec("INSERT INTO	user('username','password','phone_number','address') "+
		"VALUES(?,?,?,?)", user.Username, user.Password, user.PhoneNumber, user.Address)
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

func (p *UserDao) FindUserById(id uint) entity.User {

	var user entity.User
	row := webUtil.DB.QueryRow("SELECT * FROM user WHERE 'id' = ? ", id)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.PhoneNumber, &user.Address)
	if err != nil {
		log.Println(err)
	}
	return user
}
func (p *UserDao) Update(user entity.User) int64 {
	result, err := webUtil.DB.Exec("UPDATE user SET 'username'=?,'password'=?,'phone_number'=?,'address'=? WHERE 'id'=?’",
		user.Username, user.Password, user.PhoneNumber, user.Address, user.Id)
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
func (p *UserDao) FindUserByUsername(username string) entity.User {

	var user entity.User
	row := webUtil.DB.QueryRow("SELECT * FROM user WHERE 'username' = ? ", user)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.PhoneNumber, &user.Address)
	if err != nil {
		log.Println(err)
	}
	return user
}

func (p *UserDao) FindAllUser() []entity.User {
	rows, err := webUtil.DB.Query("SELECT * FROM user ")
	if err != nil {
		log.Println(err)
		return nil
	}
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.PhoneNumber, &user.Address)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}
	rows.Close()
	return users
}
