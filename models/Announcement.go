package models

import (
	"time"
)

// 通知 公告
type Announcement struct {
	ID string

	// 标题
	Title string

	// 内容
	Content string

	// 所属班级,一条通知可以发到多个班级
	ClassID []string

	CreateTime time.Time

	// 创建者
	CreateBy string

	UpdateTime time.Time

	// 通知是否带有图片
	HasPhoto bool
}

// 通知模板
type AnnounceTemplate struct {
	ID string

	// 模板所属分类
	TmplType int

	Title string

	ContentID string
}
