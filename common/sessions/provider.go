package sessions

import (
	"time"
)

// session的数据源
type Provider interface {
	//实现Session的初始化，操作成功则返回此新的Session变量
	SessionInit(sid string) (Session, error)

	//SessionRead函数返回sid所代表的Session变量；
	//如果不存在，那么将以sid为参数调用SessionInit函数创建并返回一个新的Session变量
	SessionRead(sid string) (Session, error)

	// 销毁sid对应的Session变量
	SessionDestroy(sid string) error

	// SessionGC根据maxLifeTime来删除过期的数据
	SessionGC(maxLifeTime time.Duration)
}
