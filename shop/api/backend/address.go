package backend

import "github.com/gogf/gf/v2/frame/g"

type AddressBase struct {
	ParentId uint   `json:"parent_id"`
	Name     string `json:"name"`
	Status   uint8  `json:"status"`
}

type AddAddressReq struct {
	g.Meta `path:"/address" tags:"Address" method:"post" summary:"add address"`
	AddressBase
}

type AddAddressRes struct {
	Id int `json:"id"`
}

type UpdateAddressReq struct {
	g.Meta `path:"/address" tags:"Address" method:"put" summary:"update address"`
	Id     int `json:"id" dc:"id"`
	AddressBase
}

type UpdateAddressRes struct {
	Id int `json:"id"`
}

type DeleteAddressReq struct {
	g.Meta `path:"/address" method:"delete" summary:"delete address"`
	Id     int `json:"id" v:"required#address id cannot be empty" dc:"address id"`
}

type DeleteAddressRes struct {
}

type PageAddressReq struct {
	g.Meta `path:"/address/page" method:"get" summary:"page address"`
	CommonPaginationReq
}

type PageAddressRes struct {
	CommonPaginationRes
}

type CityAddressReq struct {
	g.Meta `path:"/address/list" tags:"Address" method:"get" summary:"list city address"`
}

type CityAddressRes struct {
	List interface{} `json:"list" description:"address list"`
}
