package homework

import (
	"time"

	"github.com/NothingXiang/online-class/class"
)

// 作业
type Homework struct {

	// 唯一标识
	ID string `json:"id" bson:"_id"`

	// 标题
	Title string `json:"title" bson:"title"`

	// 内容
	Content string `json:"content" bson:"content"`

	// 所属班级,一条通知可以发到多个班级(未必用得上)
	Class string `json:"class" bson:"class" `

	// 作业所属的科目
	Subject class.Subject `json:"subject" bson:"subject"`

	// 通知附带的图片路径
	Photos []string `json:"photos" bson:"photo"`

	// 创建时间
	CreateTime time.Time `json:"create_time" bson:"create_time"`

	// 创建者 填用户id
	CreateBy string `json:"create_by" bson:"create_by"`

	// 更新时间
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
}
