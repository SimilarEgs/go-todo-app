package entity

type User struct {
	// validate request body with «gin-binding» fields
	Id       int    `db:"id"              json:"-"`
	Name     string `db:"name"            json:"name"                binding:"required"`
	Username string `db:"username"        json:"username"            binding:"required,alphanum"`
	Password string `db:"hashed_password" json:"password"            binding:"required,min=6"`
}
