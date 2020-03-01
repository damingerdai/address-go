package models

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" header:"username"`
	Password string `json:"password" header:"password"`
}
