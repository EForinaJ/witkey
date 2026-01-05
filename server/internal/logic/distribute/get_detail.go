package distribute

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dao_distribute "server/internal/type/distribute/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IDistribute.
func (s *sDistribute) GetDetail(ctx context.Context, id int64) (res *dao_distribute.Detail, err error) {
	witkeyId, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().UserId, ctx.Value("userId")).
		Value(dao.SysWitkey.Columns().Id)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	obj, err := dao.SysDistribute.Ctx(ctx).
		WherePri(id).
		Where(dao.SysDistribute.Columns().WitkeyId, witkeyId).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if obj.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}
	// 2. 使用结构体转换代替手动映射
	var detail *dao_distribute.Detail
	if err := gconv.Scan(obj.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.FAILD, response.CodeMsg(response.FAILD))
	}

	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, obj.GMap().Get(dao.SysDistribute.Columns().OrderId)).
		Fields(
			dao.SysOrder.Columns().Code,
			dao.SysOrder.Columns().Requirements,
			dao.SysOrder.Columns().ProductId,
			dao.SysOrder.Columns().TotalAmount,
			dao.SysOrder.Columns().Quantity,
			dao.SysOrder.Columns().Unit,
		).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Code = gconv.String(order.GMap().Get(dao.SysOrder.Columns().Code))
	detail.Quantity = gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Quantity))
	detail.Unit = gconv.String(order.GMap().Get(dao.SysOrder.Columns().Unit))
	detail.TotalAmount = gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().TotalAmount))
	detail.Requirements = gconv.String(order.GMap().Get(dao.SysOrder.Columns().Requirements))

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

	detail.Product = &dao_distribute.Product{
		Id:       gconv.Int64(order.GMap().Get(dao.SysOrder.Columns().ProductId)),
		Name:     gconv.String(product.GMap().Get(dao.SysProduct.Columns().Name)),
		Pic:      gconv.String(product.GMap().Get(dao.SysProduct.Columns().Pic)),
		Game:     game.String(),
		Category: category.String(),
	}

	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusSettlemented || gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().Status)) == consts.DistributeStatusSettlementing {
		settlementObj, err := dao.SysSettlement.Ctx(ctx).
			Where(dao.SysSettlement.Columns().OrderId, obj.GMap().Get(dao.SysDistribute.Columns().OrderId)).
			Where(dao.SysSettlement.Columns().WitkeyId, witkeyId).One()
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		var settlementDetail *dao_distribute.Settlement
		err = settlementObj.Struct(&settlementDetail)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		detail.Settlement = settlementDetail
	}

	return detail, nil
}
