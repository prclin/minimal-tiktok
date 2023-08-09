/*
Package response 存放用来映射响应数据的结构体
*/
package response

type Douyin_user_register_response struct {
	//状态代码
	StatusCode int32 `json:"status_code"`
	//状态信息
	StatusMsg string `json:"status_msg,omitempty"`
	UserId    int64  `json:"user_id"`
	Token     string `json:"token"`
}
