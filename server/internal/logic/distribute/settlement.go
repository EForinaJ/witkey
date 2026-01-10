package distribute

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_distribute "server/internal/type/distribute/dto"
	utils_code "server/internal/utils/code"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/shopspring/decimal"
)

// Settlement implements service.IDistribute.
func (s *sDistribute) Settlement(ctx context.Context, req *dto_distribute.Settlement) (err error) {
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
	witkey, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, ctx.Value("userId")).
		Fields(
			dao.SysWitkey.Columns().Id,
			dao.SysWitkey.Columns().TitleId,
		).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	obj, err := tx.Model(dao.SysDistribute.Table()).
		Where(dao.SysDistribute.Columns().Id, req.Id).
		Where(dao.SysDistribute.Columns().WitkeyId, witkey.GMap().Get(dao.SysWitkey.Columns().Id)).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	order, err := tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, obj.GMap().Get(dao.SysDistribute.Columns().OrderId)).
		Fields(dao.SysOrder.Columns().TotalAmount, dao.SysOrder.Columns().WitkeyCount).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	// 获取头衔费率
	servicePercent, err := tx.Model(dao.SysTitle.Table()).
		Where(dao.SysTitle.Columns().Id, witkey.GMap().Get(dao.SysWitkey.Columns().TitleId)).
		Value(dao.SysTitle.Columns().ServicePercent)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	// 自带队伍结算
	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Type)) == consts.DistributeTypeTeam {

		totalAmount := decimal.NewFromFloat(gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().TotalAmount)))
		percent := decimal.NewFromFloat(servicePercent.Float64()).Div(decimal.NewFromFloat(100))
		commission := totalAmount.Mul(percent)
		serviceCharge := totalAmount.Sub(commission)

		_, err = tx.Model(dao.SysSettlement.Table()).Data(g.Map{
			dao.SysSettlement.Columns().WitkeyId:      witkey.GMap().Get(dao.SysWitkey.Columns().Id),
			dao.SysSettlement.Columns().OrderId:       obj.GMap().Get(dao.SysDistribute.Columns().OrderId),
			dao.SysSettlement.Columns().DistributeId:  req.Id,
			dao.SysSettlement.Columns().Amount:        totalAmount,
			dao.SysSettlement.Columns().Commission:    commission,
			dao.SysSettlement.Columns().ServiceCharge: serviceCharge,
			dao.SysSettlement.Columns().Images:        req.Images,
			dao.SysSettlement.Columns().Code:          utils_code.GetCode(ctx, "SL"),
			dao.SysSettlement.Columns().Status:        consts.StatusApply,
			dao.SysSettlement.Columns().CreateTime:    gtime.Now(),
		}).Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}
	}

	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Type)) == consts.DistributeTypeSelf {

		totalAmount := decimal.NewFromFloat(gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().TotalAmount))).
			Sub(decimal.NewFromFloat(gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().WitkeyCount))))
		percent := decimal.NewFromFloat(servicePercent.Float64()).Div(decimal.NewFromFloat(100))
		commission := totalAmount.Mul(percent)
		serviceCharge := totalAmount.Sub(commission)

		_, err = tx.Model(dao.SysSettlement.Table()).Data(g.Map{
			dao.SysSettlement.Columns().WitkeyId:      witkey.GMap().Get(dao.SysWitkey.Columns().Id),
			dao.SysSettlement.Columns().OrderId:       obj.GMap().Get(dao.SysDistribute.Columns().OrderId),
			dao.SysSettlement.Columns().DistributeId:  req.Id,
			dao.SysSettlement.Columns().Amount:        totalAmount,
			dao.SysSettlement.Columns().Commission:    commission,
			dao.SysSettlement.Columns().ServiceCharge: serviceCharge,
			dao.SysSettlement.Columns().Images:        req.Images,
			dao.SysSettlement.Columns().Code:          utils_code.GetCode(ctx, "SL"),
			dao.SysSettlement.Columns().Status:        consts.StatusApply,
			dao.SysSettlement.Columns().CreateTime:    gtime.Now(),
		}).Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}
	}

	_, err = tx.Model(dao.SysDistribute.Table()).
		Where(dao.SysDistribute.Columns().Id, req.Id).
		Where(dao.SysDistribute.Columns().WitkeyId, witkey.GMap().Get(dao.SysWitkey.Columns().Id)).
		Data(g.Map{
			dao.SysDistribute.Columns().Status: consts.DistributeStatusSettlementing,
		}).Update()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	return
}
