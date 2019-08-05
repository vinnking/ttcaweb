package models

// 文章表，存储完整的文章信息 并映射文章提交的数据字段
type Article struct {
	Id       int    `form:"-"`
	Cid      int    `form:"cid"`
	Count    int    `form:"-"`
	Type     int    `form:"-"`    // 类型自定义 文章和页面，默认0为文章，1为页面
	Name     string `form:"name"` // 自定义链接名称
	Title    string `form:"title" orm:"size(64)"`
	Author   string `form:"author"`
	Image    string `form:"image"`
	Tags     string `form:"tags" orm:"size(64)"`
	Renew    string `form:"-" orm:"size(20)"`
	Content  string `form:"content" orm:"type(text)"`
	Summary  string `form:"summary"`
	Status   string `form:"status"`
	Password string `form:"password"`
	Link     string `form:"-"`
}
