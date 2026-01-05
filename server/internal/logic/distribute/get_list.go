package distribute

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_distribute "server/internal/type/distribute/dao"
	dto_distribute "server/internal/type/distribute/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.Idistribute.
func (s *sDistribute) GetList(ctx context.Context, req *dto_distribute.Query) (total int, res []*dao_distribute.List, err error) {
	witkeyId, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().UserId, ctx.Value("userId")).Value(dao.SysWitkey.Columns().Id)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	m := dao.SysDistribute.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysDistribute.Columns().WitkeyId, witkeyId).
		OrderDesc(dao.SysDistribute.Columns().CreateTime)
	if req.Code != "" {
		orderId, err := dao.SysOrder.Ctx(ctx).Where(dao.SysOrder.Columns().Code, req.Code).Value(dao.SysOrder.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		m = m.Where(dao.SysDistribute.Columns().OrderId, orderId)
	}
	if req.Status != 0 {
		m = m.Where(dao.SysDistribute.Columns().Status, req.Status)
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysDistribute
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_distribute.List, len(list))
	for i, v := range list {
		var entity *dao_distribute.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
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
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
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
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		game, err := dao.SysGame.Ctx(ctx).
			Where(dao.SysGame.Columns().Id,
				product.GMap().Get(dao.SysProduct.Columns().GameId)).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		category, err := dao.SysCategory.Ctx(ctx).
			Where(dao.SysCategory.Columns().Id,
				product.GMap().Get(dao.SysProduct.Columns().CategoryId)).
			Value(dao.SysCategory.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		entity.Product = &dao_distribute.Product{
			Id:       gconv.Int64(order.GMap().Get(dao.SysOrder.Columns().ProductId)),
			Name:     gconv.String(product.GMap().Get(dao.SysProduct.Columns().Name)),
			Pic:      gconv.String(product.GMap().Get(dao.SysProduct.Columns().Pic)),
			Game:     game.String(),
			Category: category.String(),
		}

		res[i] = entity
	}

	return
}
