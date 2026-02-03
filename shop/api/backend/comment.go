package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CommentDeleteReq struct {
	g.Meta `path:"/comment/delete" tags:"Comment" method:"delete" summary:"delete a comment"`
	Id     uint `json:"comment_id" v:"min:1#select a comment to delete" dc:"comment's id"`
}

type CommentDeleteRes struct {
}

type CommentGetListReq struct {
	g.Meta `path:"/comment/list" tags:"Comment" method:"get" summary:"get comment list"`
	Type   uint8 `json:"type" dc:"comments' type: 1comment 2article"`
}

type CommentGetListRes struct {
	List  interface{} `json:"list" description:"comment list"`
	Page  int         `json:"page" description:"page number"`
	Size  int         `json:"size" description:"size of page"`
	Total int         `json:"total" description:"total amount"`
}
