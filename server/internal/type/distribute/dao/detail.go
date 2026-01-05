package dao_distribute

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id           int64       `json:"id" dc:"ID"`
	Code         string      `json:"code" dc:"订单号"`
	Product      *Product    `json:"product" dc:"订单商品"`
	Quantity     int         `json:"quantity" dc:"数量"`
	Type         int         `json:"type" dc:"类型"`
	Unit         string      `json:"unit" dc:"单位"`
	TotalAmount  float64     `json:"totalAmount" dc:"订单总额"`
	Status       int         `json:"status" dc:"服务状态"`
	Reason       string      `json:"reason" dc:"取消原因"`
	Requirements string      `json:"requirements" dc:"订单要求"`
	StartTime    *gtime.Time `json:"startTime" dc:"开始时间"`
	FinishTime   *gtime.Time `json:"finishTime" dc:"结束时间"`
	Settlement   *Settlement `json:"settlement" dc:"报单结算"`
	CreateTime   *gtime.Time `json:"createTime" dc:"派单时间"`
}
