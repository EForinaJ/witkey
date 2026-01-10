package dto_account

type Edit struct {
	Name        string   `p:"name" v:"max-length:8#名称长度不能超过8"  dc:"名称"`
	Address     []string `p:"address"  dc:"地址"`
	Sex         int      `p:"sex"  dc:"性别"`
	Avatar      string   `p:"avatar" dc:"头像"`
	Description string   `p:"description"  dc:"描述"`
	Birthday    int64    `json:"birthday" dc:"生日"`
	Album       []string `json:"album" dc:"相册"`
}
