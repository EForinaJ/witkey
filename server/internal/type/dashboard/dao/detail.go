package dao_dashboard

import dao_distribute "server/internal/type/distribute/dao"

type Detail struct {
	ComparedToYesterday float64                `json:"comparedToYesterday" dc:"同比昨日"`
	TodayCommission     float64                `json:"todayCommission" dc:"今日收益"`
	DistributeList      []*dao_distribute.List `json:"distributeList" dc:"近期接单记录"`
}
