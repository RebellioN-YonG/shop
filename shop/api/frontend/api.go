package frontend

type CommonPaginationReq struct {
	Page int `json:"page" in:"query" d:"1" v:"min:0#page id error" dc:"page, default 1"`
	Size int `json:"size" in:"query" d:"10" v:"max:50#max size 50" dc:"size, max 50"`
}

type CommonPaginationRes struct {
	List  interface{} `json:"list" dc:"list of records"`
	Total int         `json:"total" dc:"total number of records"`
	Page  int         `json:"page" dc:"page number"`
	Size  int         `json:"size" dc:"size of pages' quantity"`
}
