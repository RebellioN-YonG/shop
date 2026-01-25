package backend

import "github.com/gogf/gf/v2/frame/g"

type RotationReq struct {
	g.Meta `path:"backend/rotation/add" tags: "Rotation" method:"get" summary:"ur 1st rotation api"`
}

type rotationRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
