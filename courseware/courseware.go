package courseware

import (
	"time"
)

// 课件资源
type CoursewareInfo struct {
	// 唯一id
	ID string `json:"id" bson:"_id"`

	// 文件名
	Name string `json:"name" bson:"name"`

	// 所属班级
	Class string `json:"class" bson:"class"`

	// 创建时间
	CreateTime time.Time `json:"create_time" bson:"create_time"`

	//创建人
	CreateBy string `json:"create_by" bson:"create_by"`

	// 存储路径
	Path string `json:"path" bson:"path"`
}
