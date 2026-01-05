package distribute

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_distribute "server/internal/type/distribute/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// CheckSettlement implements service.IDistribute.
func (s *sDistribute) CheckSettlement(ctx context.Context, req *dto_distribute.Settlement) (err error) {
	witkeyId, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().UserId, ctx.Value("userId")).
		Value(dao.SysWitkey.Columns().Id)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	obj, err := dao.SysDistribute.Ctx(ctx).
		Where(dao.SysDistribute.Columns().Id, req.Id).
		Where(dao.SysDistribute.Columns().WitkeyId, witkeyId).
		Fields(dao.SysDistribute.Columns().Status, dao.SysDistribute.Columns().OrderId).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusCancel {
		return utils_error.Err(response.FAILD, "派单已取消，无法报单结算")
	}
	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusPending {
		return utils_error.Err(response.FAILD, "派单待服务，无法报单结算")
	}
	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusSettlemented {
		return utils_error.Err(response.FAILD, "派单已结算，无法报单结算")
	}
	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusSettlementing {
		return utils_error.Err(response.FAILD, "派单结算中，无法报单结算")
	}
	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusInProgress {
		return utils_error.Err(response.FAILD, "派单进行中，无法报单结算")
	}

	settlementExist, err := dao.SysSettlement.Ctx(ctx).
		Where(dao.SysSettlement.Columns().WitkeyId, witkeyId).
		Where(dao.SysSettlement.Columns().OrderId, obj.GMap().Get(dao.SysDistribute.Columns().OrderId)).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if settlementExist {
		return utils_error.Err(response.FAILD, "已经提交过报单结算，请勿重复提交")
	}

	aftersalesStatus, err := dao.SysAftersales.Ctx(ctx).
		Where(dao.SysAftersales.Columns().OrderId, obj.GMap().Get(dao.SysDistribute.Columns().OrderId)).
		Value(dao.SysAftersales.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if aftersalesStatus.Int() == consts.StatusApply {
		return utils_error.Err(response.FAILD, "派单提交售后工单未处理，无法报单结算")
	}
	return
}
