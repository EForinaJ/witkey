package distribute

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// CheckComplete implements service.IDistribute.
func (s *sDistribute) CheckComplete(ctx context.Context, id int64) (err error) {

	obj, err := dao.SysDistribute.Ctx(ctx).
		Where(dao.SysDistribute.Columns().Id, id).
		Where(dao.SysDistribute.Columns().WitkeyId, ctx.Value("userId")).
		Fields(dao.SysDistribute.Columns().Status).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusCancel {
		return utils_error.Err(response.FAILD, "派单已取消，无法完成")
	}
	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusPending {
		return utils_error.Err(response.FAILD, "派单待服务，无法完成")
	}
	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusSettlemented {
		return utils_error.Err(response.FAILD, "派单已结算，无法完成")
	}
	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusSettlementing {
		return utils_error.Err(response.FAILD, "派单结算中，无法完成")
	}
	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusComplete {
		return utils_error.Err(response.FAILD, "派单已完成，无法完成")
	}
	return
}
