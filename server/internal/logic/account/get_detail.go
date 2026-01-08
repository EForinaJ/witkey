package account

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
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
	err = dao.SysUser.Ctx(ctx).Fields(dao.SysUser.Columns().Id,
		dao.SysUser.Columns().Name,
		dao.SysUser.Columns().Address,
		dao.SysUser.Columns().Sex,
		dao.SysUser.Columns().Id,
		dao.SysUser.Columns().LoginIp,
		dao.SysUser.Columns().LoginTime,
		dao.SysUser.Columns().Avatar).
		Where(dao.SysUser.Columns().Id, userId).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var witkey *dao_account.Witkey
	witkeyObj, err := dao.SysWitkey.Ctx(ctx).Where(dao.SysWitkey.Columns().UserId, userId).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	err = witkeyObj.Struct(&witkey)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	title, err := dao.SysTitle.Ctx(ctx).
		Where(dao.SysTitle.Columns().Id, witkeyObj.GMap().Get(dao.SysWitkey.Columns().TitleId)).
		Value(dao.SysTitle.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	witkey.Title = title.String()

	game, err := dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Id, witkeyObj.GMap().Get(dao.SysWitkey.Columns().GameId)).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	witkey.Game = game.String()

	res.Witkey = witkey
	err = g.Redis().SetEX(ctx, consts.Account+gconv.String(userId), res, 600)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR, response.CodeMsg(response.CACHE_SAVE_ERROR))
	}
	return
}
