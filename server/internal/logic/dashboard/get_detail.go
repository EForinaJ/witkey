package dashboard

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_dashboard "server/internal/type/dashboard/dao"
	dao_distribute "server/internal/type/distribute/dao"
	utils_common "server/internal/utils/common"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IDashboard.
func (s *sDashborad) GetDetail(ctx context.Context) (res *dao_dashboard.Detail, err error) {

	// 获取昨天收益
	yesterdayCommission, err := dao.SysSettlement.Ctx(ctx).
		Where(dao.SysSettlement.Columns().WitkeyId, ctx.Value("userId")).
		Where(dao.SysSettlement.Columns().Status, consts.StatusSuccess).
		Where("DATE(settlement_time) = ?", gtime.Now().AddDate(0, 0, -1).Format("Y-m-d")).
		Sum("commission")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	todayCommission, err := dao.SysSettlement.Ctx(ctx).
		Where(dao.SysSettlement.Columns().WitkeyId, ctx.Value("userId")).
		Where(dao.SysSettlement.Columns().Status, consts.StatusSuccess).
		Where("DATE(settlement_time) = ?", gtime.Now().Format("Y-m-d")).
		Sum("commission")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	calculate := utils_common.CalculateChangePercent(todayCommission, yesterdayCommission)

	detail := dao_dashboard.Detail{
		ComparedToYesterday: calculate,
		TodayCommission:     todayCommission,
	}

	// 获取派单列表
	var list []*entity.SysDistribute
	err = dao.SysDistribute.Ctx(ctx).
		Where(dao.SysDistribute.Columns().WitkeyId, ctx.Value("userId")).Page(1, 10).
		OrderDesc(dao.SysDistribute.Columns().CreateTime).Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	distributeList := make([]*dao_distribute.List, len(list))
	for i, v := range list {
		var entity *dao_distribute.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		order, err := dao.SysOrder.Ctx(ctx).Where(dao.SysOrder.Columns().Id, v.OrderId).
			Fields(dao.SysOrder.Columns().ProductId,
				dao.SysOrder.Columns().Code,
				dao.SysOrder.Columns().UnitPrice,
				dao.SysOrder.Columns().Quantity,
				dao.SysOrder.Columns().TotalAmount,
				dao.SysOrder.Columns().Unit,
			).
			One()
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Quantity = gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Quantity))
		entity.Unit = gconv.String(order.GMap().Get(dao.SysOrder.Columns().Unit))
		entity.Code = gconv.String(order.GMap().Get(dao.SysOrder.Columns().Code))
		entity.TotalAmount = gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().TotalAmount))
		// 商品内容
		product, err := dao.SysProduct.Ctx(ctx).
			Where(dao.SysProduct.Columns().Id, order.GMap().Get(dao.SysOrder.Columns().ProductId)).
			Fields(
				dao.SysProduct.Columns().Name,
				dao.SysProduct.Columns().Pic,
				dao.SysProduct.Columns().GameId,
				dao.SysProduct.Columns().CategoryId,
			).
			One()
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		game, err := dao.SysGame.Ctx(ctx).
			Where(dao.SysGame.Columns().Id,
				product.GMap().Get(dao.SysProduct.Columns().GameId)).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		category, err := dao.SysCategory.Ctx(ctx).
			Where(dao.SysCategory.Columns().Id,
				product.GMap().Get(dao.SysProduct.Columns().CategoryId)).
			Value(dao.SysCategory.Columns().Name)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		entity.Product = &dao_distribute.Product{
			Id:       gconv.Int64(order.GMap().Get(dao.SysOrder.Columns().ProductId)),
			Name:     gconv.String(product.GMap().Get(dao.SysProduct.Columns().Name)),
			Pic:      gconv.String(product.GMap().Get(dao.SysProduct.Columns().Pic)),
			Game:     game.String(),
			Category: category.String(),
		}

		distributeList[i] = entity
	}
	detail.DistributeList = distributeList

	return &detail, nil
}
