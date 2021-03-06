package service

import (
	"github.com/NothingXiang/online-class/notice"
)

type NoticeService interface {
	// 根据通知ID获取通知
	GetNotice(noticeID string) (*notice.Notice, error)

	// 获取通知的已读列表
	GetReadList(noticeID string) ([]string, error)

	// 分页获取班级公告
	GetNoticeByClass(classID string, page, limit int) ([]*notice.Notice, error)

	//	发布公告
	CreateNotice(n *notice.Notice) error

	// 移除通知
	RemoveNotice(noticeID string) error

	//	编辑公告
	UpdateNotice(update *notice.Notice) error

	// 获取某个类别的通知模板
	GetNoticeTemplate(noticeType string) ([]*notice.Template, error)

	// 往通知的已读列表添加成员（只添加学生）
	AddNoticeReadList(noticeID, userID string) error

	// 批量创建通知模板
	//BatchCreateTemplate(tmpls []*notice.Template) error
}
