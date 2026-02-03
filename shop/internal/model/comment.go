package model

import (
	"shop/internal/model/do"

	"github.com/gogf/gf/v2/os/gtime"
)

// add comment input
type AddCommentInput struct {
	UserId   uint
	ObjectId uint
	Type     uint8
	ParentId uint
	Content  string
}

// add comment output
type AddCommentOutput struct {
	Id uint
}

// delete comment input
type CommentDeleteInput struct {
	Id uint
}

// delete comment output
type CommentDeleteOutput struct {
	Id uint
}

// comment list input
type CommentListInput struct {
	Page     int
	Size     int
	ObjectId uint
	Type     uint8
}

// comment list output
type CommentListOutput struct {
	List  []CommentListOutputItem `json:"list" description:"list"`
	Page  int                     `json:"page" description:"page"`
	Size  int                     `json:"size" description:"size"`
	Total int                     `json:"total" description:"total"`
}

// comment list output item
type CommentListOutputItem struct {
	Id        int         `json:"id" description:"comment id"`
	UserId    int         `json:"userId" description:"user id"`
	ObjectId  int         `json:"objectId" description:"object id (goods id or article id)"`
	Type      int         `json:"type" description:"comment type: 1 goods 2 article"`
	ParentId  uint        `json:"parentId" description:"parent comment id"`
	Content   string      `json:"content" description:"comment content"`
	CreatedAt *gtime.Time `json:"created_at" description:"create time"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"update time"`
	// todo:
	// Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	// Article   ArticleItem `json:"article" orm:"with:id=object_id"`
}

// comment base struct for nested usage
type CommentBase struct {
	do.CommentInfo
	// todo:
	User    UserInfoBase `json:"user" orm:"with:id=user_id"`
	ReplyTo string       `json:"reply_to" `
}
