package distribute

import (
	"context"

	v1 "server/api/distribute/v1"
	"server/internal/service"
)

func (c *ControllerV1) Start(ctx context.Context, req *v1.StartReq) (res *v1.StartRes, err error) {
	err = service.Distribute().CheckStart(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = service.Distribute().Start(ctx, req.Id)
	return
}
