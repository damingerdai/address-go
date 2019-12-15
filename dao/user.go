package dao

import (
	"damingerdai/address/models"
	"errors"
)

func CreateUser(id int64, username, password string) (err error) {
	sql := "INSERT user SET id = ? ,name = ?, password = ?"
	stmt, err := conn.Prepare(sql)
	if err != nil {
		return
	}
	res, err := stmt.Exec(id, username, password)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect == 0 {
		err = errors.New("fail to create a user")
	}
	return
}

func GetUserById(id int64) (*models.User, error) {
	rows := conn.QueryRow("SELECT name, password FROM user WHERE id = ?", id)
	var name, password string
	err := rows.Scan(&name, &password)
	if err != nil {
		return nil, err
	}
	user := models.User{id, name, password}
	return &user, nil
}
