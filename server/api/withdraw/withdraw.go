// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package withdraw

import (
	"context"

	"server/api/withdraw/v1"
)

type IWithdrawV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
}
