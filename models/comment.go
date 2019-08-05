package models

// 评论列表
type Comment struct {
	Id      int    `form:"-"`       //id
	Aid     int    `form:"aid"`     //所属文章id
	Status  int    `form:"-"`       // 评论的状态 0：no，1：yes
	Reply   int    `form:"-"`       // 对评论的回复
	Content string `form:"content"` // 回复内容
	Date    string `form:"-"`       // 回复时间
	Email   string `form:"email"`   // 邮箱
	Name    string `form:"name"`    // 名称
	Link    string `form:"link"`    // 链接
}
