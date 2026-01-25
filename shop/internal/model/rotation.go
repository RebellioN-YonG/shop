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
