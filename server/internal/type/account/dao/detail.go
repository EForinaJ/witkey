package dao_account

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id        int16       `json:"id" dc:"账户id"`
	Name      string      `json:"name" dc:"账户昵称"`
	Avatar    string      `json:"avatar" dc:"头像"`
	Sex       int         `json:"sex" dc:"账户性别"`
	Address   []string    `json:"address" dc:"账户地址"`
	Witkey    *Witkey     `json:"witkey" dc:"账户描述"`
	LoginTime *gtime.Time `json:"loginTime" dc:"登录时间"`
	LoginIp   string      `json:"LoginIp" dc:"登录地址"`
}
