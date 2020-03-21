package store

import (
	"github.com/NothingXiang/online-class/notice"
)

type NoticeStore interface {

	// 分页获取班级公告
	GetNoticeByClass(classID string, page, limit int) ([]*notice.Notice, error)

	//	发布公告
	CreateNotice(n *notice.Notice) error

	// 移除通知 返回班级id和error
	RemoveNotice(noticeID string) (string,error)

	//	编辑公告
	UpdateNotice(update *notice.Notice) error

	// 获取某个类别的通知模板
	GetNoticeTemplate(noticeType string) ([]*notice.Template, error)

	// 批量创建通知模板
	//BatchCreateTemplate(tmpls []*notice.Template) error
}
