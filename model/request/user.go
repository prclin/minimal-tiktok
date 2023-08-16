package request

// UserEnrollRequest 用户注册登录请求参数
type UserEnrollRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
