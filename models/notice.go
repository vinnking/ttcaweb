package models

//  公告表
type Notice struct {
	Id      int    // 自增唯一ID
	Content string // 内容
	Date    string // 发布时间
	Url     string // 定向URL
}
