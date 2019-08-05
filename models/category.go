package models

// category分类表，慢慢优化利用
type Category struct {
	Id          int    // 分类方法ID
	Name        string // 分类名称
	Key         string // 分类标识
	Method      string // 分类方法 文章分类article 上传分类upload
	Description string // 分类描述
	Parent      string // 所属父分类方法ID
}

type CategoryData struct {
	IsNil bool // Check this category is null
	Msg   string
	Info  []Category
	List  []Article // this article include in this category content
}
