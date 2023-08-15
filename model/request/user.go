package request

// UserRegisterRequest 用户注册请求参数
type UserRegisterRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
