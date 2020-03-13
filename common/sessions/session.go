package sessions

import (
	"sync"
)

// session 的抽象接口
type Session interface {
	Set(key string, value interface{}) error //设置session value
	Get(key string) (interface{}, bool)      //获取session value
	Delete(key string) error                 //删除session
	SessionID() string                       //back current sessionID
}

// 包内提供的最基本的储存key-value的session
type BaseSess struct {
	// session 唯一标识
	sid string

	// 保护store的读写锁
	sync.RWMutex

	// 存储session value
	store map[string]interface{}
}

func (K *BaseSess) Set(key string, value interface{}) error {
	K.Lock()
	defer K.Unlock()

	K.store[key] = value
	return nil
}

func (K *BaseSess) Get(key string) (r interface{}, exist bool) {
	K.RLock()
	defer K.RUnlock()

	r, exist = K.store[key]
	return
}

func (K *BaseSess) Delete(key string) error {
	K.Lock()
	defer K.Unlock()

	delete(K.store, key)

	return nil
}

func (K *BaseSess) SessionID() string {
	return K.sid
}
