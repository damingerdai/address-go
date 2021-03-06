package dao

import (
	"damingerdai/address/internal/models"
	"errors"
	"github.com/jmoiron/sqlx"
	"strings"
)

func CreateUser(trx *sqlx.Tx, id int64, username, password string) (err error) {
	sql := "INSERT user SET id = ? ,name = ?, password = ?"
	stmt, err := trx.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
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

func GetUserById(trx *sqlx.Tx, id int64) (*models.User, error) {
	rows := trx.QueryRow("SELECT name, password FROM user WHERE id = ?", id)
	var name, password string
	err := rows.Scan(&name, &password)
	if err != nil {
		return nil, err
	}
	user := models.User{id, name, password}
	return &user, nil
}

func ListUsers(trx *sqlx.Tx) (*[]models.User, error) {
	rows, err := trx.Query("SELECT id, name, password FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]models.User, 0)
	for rows.Next() {
		var id int64
		var name, password string
		err = rows.Scan(&id, &name, &password)
		if err != nil {
			panic(err.Error())
		}
		user := models.User{id, name, password}
		users = append(users, user)
	}

	return &users, nil
}

func HasUser(trx *sqlx.Tx, user *models.User) (bool, error) {
	sql := "SELECT count(id) FROM user where "
	conditions := make([]string, 0, 3)
	params := make([]interface{}, 0, 3)
	if user.ID > 0 {
		conditions = append(conditions, " id = ? ")
		params = append(params, user.ID)
	}
	if len(user.Username) > 0 {
		conditions = append(conditions, " name = ? ")
		params = append(params, user.Username)
	}
	if len(user.Password) > 0 {
		conditions = append(conditions, " password = ? ")
		params = append(params, user.Password)
	}
	rows := trx.QueryRow(sql+sliceConditions(&conditions), params...)
	var count int
	err := rows.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil

}

func sliceConditions(conditions *[]string) string {
	var b strings.Builder
	for i, v := range *conditions {
		if i != 0 {
			b.WriteString(" and ")
		}
		b.WriteString(v)
	}

	return b.String()
}
