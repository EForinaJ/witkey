package dao_account

type Witkey struct {
	Id         int16    `json:"id" dc:"账户id"`
	Name       string   `json:"name" dc:"账户昵称"`
	Title      string   `json:"title" dc:"头衔"`
	Game       string   `json:"game" dc:"游戏领域"`
	Commission float64  `json:"commission" dc:"佣金"`
	Album      []string `json:"album" dc:"相册"`
	Rate       int      `json:"rate" dc:"评分"`
}
