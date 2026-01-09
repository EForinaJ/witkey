package withdraw

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_withdraw "server/internal/type/withdraw/dto"
	utils_code "server/internal/utils/code"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// Create implements service.IWithdraw.
func (s *sWithdraw) Create(ctx context.Context, req *dto_withdraw.Create) (err error) {
	witkeyId, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().UserId, ctx.Value("userId")).Value(dao.SysWitkey.Columns().Id)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	withdrawSetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WithdrawSetting).Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	withdrawJson, err := gjson.DecodeToJson(withdrawSetting)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	withDrawPercent := withdrawJson.Get("withDrawPercent").Float64()
	percent := decimal.NewFromFloat(withdrawJson.Get("withDrawPercent").Float64()).Div(decimal.NewFromFloat(100))
	amount := decimal.NewFromFloat(req.Amount)
	settledAmount := amount
	serviceFee := decimal.NewFromFloat(0)
	if withDrawPercent != 0 {
		serviceFee = amount.Mul(percent)
		settledAmount = amount.Sub(serviceFee)
	}
	_, err = dao.SysWithdraw.Ctx(ctx).Data(g.Map{
		dao.SysWithdraw.Columns().WitkeyId:      witkeyId,
		dao.SysWithdraw.Columns().Amount:        amount,
		dao.SysWithdraw.Columns().SettledAmount: settledAmount,
		dao.SysWithdraw.Columns().ServiceFee:    serviceFee,
		dao.SysWithdraw.Columns().Status:        consts.StatusApply,
		dao.SysWithdraw.Columns().CreateTime:    gtime.Now(),
		dao.SysWithdraw.Columns().Name:          req.Name,
		dao.SysWithdraw.Columns().Number:        req.Number,
		dao.SysWithdraw.Columns().Type:          req.Type,
		dao.SysWithdraw.Columns().Code:          utils_code.GetCode(ctx, consts.TX),
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	return
}
