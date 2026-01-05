// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package distribute

import (
	"context"

	"server/api/distribute/v1"
)

type IDistributeV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	Start(ctx context.Context, req *v1.StartReq) (res *v1.StartRes, err error)
	Complete(ctx context.Context, req *v1.CompleteReq) (res *v1.CompleteRes, err error)
	Settlement(ctx context.Context, req *v1.SettlementReq) (res *v1.SettlementRes, err error)
}
