package model

//User 结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Number   string
}
type Judge struct {
	IsLogin bool
}