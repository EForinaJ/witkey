package auth

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/lib/jwt"
	"server/internal/model/entity"
	dto_auth "server/internal/type/auth/dto"
	utils_error "server/internal/utils/error"
	utils_lock "server/internal/utils/lock"
	"server/internal/utils/response"
	"strings"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Login implements service.IAuth.
func (s *sAuth) Login(ctx context.Context, req *dto_auth.Login) (res interface{}, err error) {
	obj, err := dao.SysWitkey.Ctx(ctx).Fields(
		dao.SysWitkey.Columns().Salt,
		dao.SysWitkey.Columns().Password,
		dao.SysWitkey.Columns().Id,
	).Where(dao.SysWitkey.Columns().Phone, req.Phone).One()

	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if obj == nil {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}
	var witkey entity.SysWitkey
	err = obj.Struct(&witkey)
	if err != nil {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}
	randPwd := consts.SYSTEMNAME + req.PassWord + witkey.Salt
	randPwd = gmd5.MustEncryptString(randPwd)

	if !strings.EqualFold(witkey.Password, randPwd) {
		// 设置密码错误次数
		errTimes, _ := utils_lock.SetCount(ctx, consts.LoginCount+req.Phone,
			consts.LoginLock+req.Phone, 1800, 5)
		having := 5 - errTimes
		if having == 0 {
			_, err = g.Redis().Del(ctx, consts.LoginCount+req.Phone)
			if err != nil {
				return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
			}
			return nil, utils_error.Err(response.LOGIN_ERROR, "账号已锁定，请30分钟后再试")
		} else {
			return nil, utils_error.Err(response.LOGIN_ERROR, "密码不正确"+",还有"+gconv.String(having)+"次之后账号将锁定")
		}
	} else {
		_, err = dao.SysWitkey.Ctx(ctx).
			Where(dao.SysWitkey.Columns().Id, witkey.Id).
			Data(g.Map{
				dao.SysWitkey.Columns().LoginIp:   g.RequestFromCtx(ctx).GetClientIp(),
				dao.SysWitkey.Columns().LoginTime: gtime.Now(),
			}).Update()
		if err != nil {
			return "", utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}
		token, err := jwt.GenToken(witkey.Id, witkey.Name)
		if err != nil {
			return "", utils_error.Err(response.AUTH_ERROR, response.CodeMsg(response.AUTH_ERROR))
		}
		res = token
	}
	return
}
