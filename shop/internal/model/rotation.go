package model

type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
}

// create rotation input
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// create rotation output
type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}

type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}

type RotationGetListInput struct {
	Page int
	Size int
	Sort int
}

type RotationGetListOutput struct {
	List  []RotationGetListOutputItem `json:"list" description:"list"`
	Total int                         `json:"total" description:"total"`
	Page  int                         `json:"page" description:"page"`
	Size  int                         `json:"size" description:"size"`
}

type RotationGetListOutputItem struct {
	Id         uint   `json:"id"`
	PicUrl     string `json:"pic_url"`
	Link       string `json:"link"`
	Sort       int    `json:"sort"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type RotationSearchInput struct {
	Key        string // search key
	Type       string // search type
	CategoryId uint   // category id
	Page       int    // page
	Size       int    // size: max 50
	Sort       int    // sort: 0 newest, 1 activity, 2 hot
}

type RotationSearchOutput struct {
	List  []RotationGetListOutput `json:"list"`
	Stats map[string]int          `json:"stats"`
	Page  int                     `json:"page"`
	Size  int                     `json:"size"`
	Total int                     `json:"total"`
}
