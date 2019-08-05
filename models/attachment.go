package models

// 附件资源记录表
type Attachment struct {
	Id      int                     // id
	Name    string                  // 名称
	Type    string `orm:"size(64)"` // 分类，独立的文件 属于article的文章
	Path    string                  // 路径
	Created string `orm:"size(10)"` // 创建时间
}
