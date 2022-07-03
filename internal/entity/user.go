package entity

type User struct {
	// validate request body with «gin-binding» fields
	Id       int    `json:"-"`
	Name     string `json:"name"     binding:"required"`
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}
