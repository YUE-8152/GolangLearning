package po

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Age      int    `json:"age" form:"age"`
	Password string `json:"password" form:"password"`
}
