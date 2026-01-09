package withdraw

import (
	"context"

	v1 "server/api/withdraw/v1"
	"server/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.Withdraw().CheckCreate(ctx, req.Create)
	if err != nil {
		return nil, err
	}
	err = service.Withdraw().Create(ctx, req.Create)
	return
}
