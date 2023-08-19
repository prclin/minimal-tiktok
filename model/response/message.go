package response

import "github.com/prclin/minimal-tiktok/model/entity"

type ChatResponse struct {
	Response
	MessageList []entity.Message `json:"message_list"`
}
