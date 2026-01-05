package dao_distribute

import "github.com/gogf/gf/v2/os/gtime"

type Settlement struct {
	Id            int64       `json:"id"     dc:"报单id"`
	Code          string      `json:"code" dc:"报单编号"`
	Amount        float64     `json:"amount" dc:"结算金额"`
	Commission    float64     `json:"commission" dc:"所得佣金"`
	ServiceCharge float64     `json:"serviceCharge" dc:"手续费'"`
	Status        int         `json:"status" dc:"状态"`
	Reason        string      `json:"reason" dc:"拒绝原因"`
	Images        []string    `json:"images" dc:"图片"`
	CreateTime    *gtime.Time `json:"createTime" dc:"申请时间"`
}
