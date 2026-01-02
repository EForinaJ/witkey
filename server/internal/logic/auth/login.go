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
	obj := dao.SysUser.Ctx(ctx).Fields(
		dao.SysUser.Columns().Salt,
		dao.SysUser.Columns().Password,
		dao.SysUser.Columns().Id,
	)
	obj = obj.Where(dao.SysUser.Columns().Phone, req.Phone)

	userObj, err := obj.One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if userObj == nil {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}
	var user entity.SysUser
	userObj.Struct(&user)
	randPwd := consts.SYSTEMNAME + req.PassWord + user.Salt
	randPwd = gmd5.MustEncryptString(randPwd)

	if !strings.EqualFold(user.Password, randPwd) {
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
		// 判断是否是威客
		exist, err := dao.SysWitkey.Ctx(ctx).
			Where(dao.SysWitkey.Columns().UserId, user.Id).
			Exist()
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		if !exist {
			return nil, utils_error.Err(response.LOGIN_ERROR, "该用户不是威客，无法登录")
		}

		_, err = dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, user.Id).
			Data(g.Map{
				dao.SysUser.Columns().LoginIp:   g.RequestFromCtx(ctx).GetClientIp(),
				dao.SysUser.Columns().LoginTime: gtime.Now(),
			}).Update()
		if err != nil {
			return "", utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}
		// witkeyId := gconv.Int64(witkey.GMap().Get(dao.SysWitkey.Columns().Id))
		// witkeyName := gconv.String(witkey.GMap().Get(dao.SysWitkey.Columns().Name))
		token, err := jwt.GenToken(user.Id, user.Name)
		if err != nil {
			return "", utils_error.Err(response.AUTH_ERROR, response.CodeMsg(response.AUTH_ERROR))
		}
		res = token
	}
	return
}
