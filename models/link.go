package models

// 链接表
type Link struct {
	Id          int    // 自增唯一ID
	Url         string // 链接URL
	Name        string // 链接名称
	Image       string // 链接图片
	Target      string // 链接打开方式
	Description string // 链接描述
	Visible     string // 是否可见（Y/N）
	Type        int    // 链接类型，普通友情链接 菜单链接
	Sort        int    // 链接排序
}
