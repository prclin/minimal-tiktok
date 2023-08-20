package response

import "github.com/prclin/minimal-tiktok/model/entity"

type CommentInfo struct {
	entity.Comment
	User UserInfo `json:"user"`
}

type CommentResponse struct {
	Response
	Comment *CommentInfo `json:"comment"`
}
