package account

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_account "server/internal/type/account/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.IAccount.
func (s *sAccount) Edit(ctx context.Context, req *dto_account.Edit) (err error) {

	_, err = dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, ctx.Value("userId")).
		Data(g.Map{
			dao.SysWitkey.Columns().Name:        req.Name,
			dao.SysWitkey.Columns().Sex:         req.Sex,
			dao.SysWitkey.Columns().Address:     req.Address,
			dao.SysWitkey.Columns().Avatar:      req.Avatar,
			dao.SysWitkey.Columns().Description: req.Description,
			dao.SysWitkey.Columns().Birthday:    gtime.NewFromTimeStamp(req.Birthday / 1000).UTC(),
			dao.SysWitkey.Columns().Album:       req.Album,
		}).Update()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	_, err = g.Redis().Del(ctx, consts.Account+gconv.String(ctx.Value("userId")))
	if err != nil {
		return utils_error.Err(response.CACHE_SAVE_ERROR, response.CodeMsg(response.CACHE_SAVE_ERROR))
	}
	return
}
