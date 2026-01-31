package model

import (
	"shop/api/backend"

	"github.com/gogf/gf/v2/frame/g"
)

type AddressBase struct {
	ParentId uint   `json:"parent_id" orm:"pid"    dc:"父级id"`
	Name     string `json:"name"      orm:"name"   dc:"地址名称"`
	Status   uint8  `json:"status"    orm:"status" dc:"状态"`
}

type AddAddressInput struct {
	AddressBase
}

type AddAddressOutput struct {
	Id int `json:"id"`
}

type UpdateAddressInput struct {
	Id int `json:"id"`
	AddressBase
}

type UpdateAddressOutput struct {
	Id int `json:"id"`
}

type DeleteAddressInput struct {
	Id int `json:"id"`
}

type DeleteAddressOutput struct {
}

type PageAddressInput struct {
	backend.CommonPaginationReq
}

type PageAddressOutput struct {
	backend.CommonPaginationRes
}

type CityAddressListOutput struct {
	List []CityAddressListOutputItem `json:"list" description:"address list"`
}

type CityAddressListOutputItem struct {
	g.Meta   `orm:"table:address_info"`
	Id       uint   `json:"id"`
	ParentId uint   `json:"parent_id"`
	Name     string `json:"name"`
	Status   uint8  `json:"status"`
}
