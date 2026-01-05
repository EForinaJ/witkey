package distribute

import (
	"context"

	v1 "server/api/distribute/v1"
	"server/internal/service"
)

func (c *ControllerV1) Settlement(ctx context.Context, req *v1.SettlementReq) (res *v1.SettlementRes, err error) {
	err = service.Distribute().CheckSettlement(ctx, req.Settlement)
	if err != nil {
		return nil, err
	}
	err = service.Distribute().Settlement(ctx, req.Settlement)
	return
}
