package withdraw

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_withdraw "server/internal/type/withdraw/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// CheckCreate implements service.IWithdraw.
func (s *sWithdraw) CheckCreate(ctx context.Context, req *dto_withdraw.Create) (err error) {
	witkeyId, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().UserId, ctx.Value("userId")).Value(dao.SysWitkey.Columns().Id)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	// 判断有提现记录还没审核
	exist, err := dao.SysWithdraw.Ctx(ctx).
		Where(dao.SysWithdraw.Columns().WitkeyId, witkeyId).
		Where(dao.SysWithdraw.Columns().Status, consts.StatusApply).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if exist {
		return utils_error.Err(response.FAILD, "有提现申请还未审核，无法新申请提现")
	}

	commission, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, witkeyId).
		Value(dao.SysWitkey.Columns().Commission)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if decimal.NewFromFloat(commission.Float64()).
		LessThan(decimal.NewFromFloat(req.Amount)) {
		return utils_error.Err(response.FAILD, "佣金不足，无法新申请提现")
	}

	// 获取最小提现额度
	withdrawSetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WithdrawSetting).Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	withdrawJson, err := gjson.DecodeToJson(withdrawSetting)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	minWithdraw := withdrawJson.Get("minWithdraw").Float64()

	if decimal.NewFromFloat(req.Amount).
		LessThan(decimal.NewFromFloat(minWithdraw)) {
		return utils_error.Err(response.FAILD, "未达到最小提现金额"+"最小提现金额: "+gconv.String(minWithdraw))
	}

	return
}
