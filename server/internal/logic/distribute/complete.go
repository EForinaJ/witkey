package distribute

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Complete implements service.IDistribute.
func (s *sDistribute) Complete(ctx context.Context, id int64) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_TX_ERROR, response.CodeMsg(response.DB_TX_ERROR))
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	witkeyId, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().UserId, ctx.Value("userId")).
		Value(dao.SysWitkey.Columns().Id)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	orderId, err := tx.Model(dao.SysDistribute.Table()).
		Where(dao.SysDistribute.Columns().Id, id).
		Where(dao.SysDistribute.Columns().WitkeyId, witkeyId).
		Value(dao.SysDistribute.Columns().OrderId)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	orderStatus, err := tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, orderId).
		Value(dao.SysOrder.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if orderStatus.Int() == consts.OrderStatusInProgress {
		_, err = tx.Model(dao.SysOrder.Table()).Where(dao.SysOrder.Columns().Id, orderId).Data(g.Map{
			dao.SysOrder.Columns().Status:     consts.OrderStatusComplete,
			dao.SysOrder.Columns().FinishTime: gtime.Now(),
		}).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}
	}

	_, err = tx.Model(dao.SysDistribute.Table()).
		Where(dao.SysDistribute.Columns().Id, id).Data(g.Map{
		dao.SysDistribute.Columns().Status:     consts.DistributeStatusComplete,
		dao.SysDistribute.Columns().FinishTime: gtime.Now(),
	}).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}

	return
}
