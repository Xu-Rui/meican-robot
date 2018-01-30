package model

//UserConfig 用户相关配置
type UserConfig struct {
	Username  string `validate:"email"` //账号
	Password  string `validate:"gt=0"`  //密码
	Workspace string `validate:"gt=0"`  //工作地点
	IsRandom  bool   //是否随机订餐
}
