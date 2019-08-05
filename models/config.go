package models


// 全局配置表 option=配置选项，value=配置内容
type Config struct {
	Id     int
	Option string `orm:"size(16)"` // 配置选项
	Value  string `orm:"size(32)"` // 选项内容
}