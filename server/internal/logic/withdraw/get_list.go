package withdraw

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_withdraw "server/internal/type/withdraw/dao"
	dto_withdraw "server/internal/type/withdraw/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IWithdraw.
func (s *sWithdraw) GetList(ctx context.Context, req *dto_withdraw.Query) (total int, res []*dao_withdraw.List, err error) {

	m := dao.SysWithdraw.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysWithdraw.Columns().WitkeyId, ctx.Value("userId")).
		OrderDesc(dao.SysWithdraw.Columns().CreateTime)
	if req.Code != "" {
		m = m.Where(dao.SysWithdraw.Columns().Code, req.Code)
	}
	if req.Status != 0 {
		m = m.Where(dao.SysWithdraw.Columns().Status, req.Status)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysWithdraw
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_withdraw.List, len(list))
	for i, v := range list {
		var entity *dao_withdraw.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		res[i] = entity
	}

	return
}
