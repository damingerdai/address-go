package service

import (
	"damingerdai/address/models"
	"damingerdai/address/utils"
	"fmt"
)

func CreateUser(user *models.User) (err error) {
	id, err := utils.CreateID()
	if err != nil {
		return err
	}
	user.ID = id
	fmt.Println(user)
	fmt.Printf("create a user: id = %d", user.ID)
	return nil
}
