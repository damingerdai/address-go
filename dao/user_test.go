package dao

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("./../.env")
	fmt.Println("load env file")
}

func TestGetUserById(t *testing.T) {
	ids := [2]int64{1206111486875799552, 100}
	for _, v := range ids {
		user, err := GetUserById(v)
		if err != nil && err != sql.ErrNoRows {
			t.Error(err)
		} else {
			t.Log(user)
		}
	}

}

func TestHasUser(t *testing.T) {
	users, err := ListUsers()
	if err != nil {
		t.Error(err.Error())
	}
	for _, v := range users {
		b, err := HasUser(v)
		if err != nil {
			t.Error(err.Error())
			break
		}
		if b != true {
			t.Errorf("no user id : %d", v.ID)
		}
	}
}
