package v1

import (
	dao_withdraw "server/internal/type/withdraw/dao"
	dto_withdraw "server/internal/type/withdraw/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/withdraw/list" method:"get" tags:"提现" summary:"提现列表"`
	*dto_withdraw.Query
}
type GetListRes struct {
	Total int                  `json:"total" dc:"总数"`
	List  []*dao_withdraw.List `json:"list" dc:"提现列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/withdraw/detail" method:"get" tags:"提现" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_withdraw.Detail
}

type CreateReq struct {
	g.Meta `path:"/withdraw/create" method:"post" tags:"提现" summary:"创建提现"`
	*dto_withdraw.Create
}
type CreateRes struct{}
