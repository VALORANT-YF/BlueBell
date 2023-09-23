package models

// UserRegisterParams 定义注册时的请求参数结构体
type UserRegisterParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	//binding:"required,eqfield=Password"` eqfidld = Password 判断两次密码是否相等
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// UserLoginParams 用户登录时的请求参数的结构体
type UserLoginParams struct {
	Username string `json:"username" bind:"required" form:"username"`
	Password string `json:"password" bind:"required" form:"password"`
}
