package v1

import (
	dao_distribute "server/internal/type/distribute/dao"
	dto_distribute "server/internal/type/distribute/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/distribute/list" method:"get" tags:"派单" summary:"派单列表"`
	*dto_distribute.Query
}
type GetListRes struct {
	Total int                    `json:"total" dc:"总数"`
	List  []*dao_distribute.List `json:"list" dc:"派单列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/distribute/detail" method:"get" tags:"派单" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_distribute.Detail
}

type StartReq struct {
	g.Meta `path:"/distribute/start" method:"post" tags:"派单" summary:"开始服务"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type StartRes struct{}

type CompleteReq struct {
	g.Meta `path:"/distribute/complete" method:"post" tags:"派单" summary:"完成服务"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type CompleteRes struct{}

type SettlementReq struct {
	g.Meta `path:"/distribute/settlement" method:"post" tags:"派单" summary:"报单结算"`
	*dto_distribute.Settlement
}
type SettlementRes struct{}
