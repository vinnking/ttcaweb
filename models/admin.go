package models

// 在后台显示的文章操作界面 数据结构
type DisplayArticle struct {
	Id     int
	Title  string
	Author string
	Name   string
	Tags   string
	Renew  string
	Uuid   string
}

// 在后台显示的评论列表
type DisplayComment struct {
	Id      int
	Name    string
	Content string
	Title   string
	Date    string
}

// 网站后台配置表单字段，
type SiteConfigOption struct {
	WebTitle       string `form:"WebTitle"`
	Keywords       string `form:"keywords"`
	WebUrl         string `form:"weburl"`
	Description    string `form:"description"`
	Status         string `form:"status"`
	Theme          string `form:"theme"`
	PageSize       string `form:"pagesize"`
	CopyRight      string `form:"copyright"`
	LogoUrl        string `form:"logourl"`
	SecondaryTitle string `form:"secondarytitle"`
}
