package account

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_account "server/internal/type/account/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IAccount.
func (s *sAccount) GetDetail(ctx context.Context) (res *dao_account.Detail, err error) {
	userId := ctx.Value("userId")
	options, err := g.Redis().Get(ctx, consts.Account+gconv.String(userId))
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
	}

	if !options.IsEmpty() {
		err = options.Scan(&res)
		if err != nil {
			return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
		}
		return
	}
	var witkey *entity.SysWitkey
	err = dao.SysWitkey.Ctx(ctx).Fields(dao.SysWitkey.Columns().Id,
		dao.SysWitkey.Columns().Name,
		dao.SysWitkey.Columns().Address,
		dao.SysWitkey.Columns().Sex,
		dao.SysWitkey.Columns().Id,
		dao.SysWitkey.Columns().LoginIp,
		dao.SysWitkey.Columns().LoginTime,
		dao.SysWitkey.Columns().Avatar).
		Where(dao.SysWitkey.Columns().Id, userId).Scan(&witkey)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	var detail *dao_account.Detail
	err = gconv.Scan(witkey, &detail)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	title, err := dao.SysTitle.Ctx(ctx).
		Where(dao.SysTitle.Columns().Id, witkey.TitleId).
		Value(dao.SysTitle.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Title = title.String()

	game, err := dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Id, witkey.GameId).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Game = game.String()

	err = g.Redis().SetEX(ctx, consts.Account+gconv.String(userId), detail, 600)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR, response.CodeMsg(response.CACHE_SAVE_ERROR))
	}
	return detail, nil
}
