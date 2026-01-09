package dto_withdraw

type Create struct {
	Amount float64 `p:"amount"  v:"required#请输入金额"   dc:"金额"`
	Name   string  `p:"name"  v:"required#请输入账户名称"   dc:"账户名称"`
	Number string  `p:"number"  v:"required#请输入账户号码"   dc:"账户号码"`
	Type   int     `p:"type"  v:"required|in:1,2#请选择提现平台|平台类型仅1或2"   dc:"提现平台"`
}
