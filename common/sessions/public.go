/*
 todo:ignore this package : unfinished
*/
package sessions

import (
	"log"
	"time"
)

var (
	// 全局唯一的一个session管理器,保外所使用的功能都要通过该管理器调用
	GlobalManager *Manager

	// 统一的session数据提供者的集合
	provides map[string]Provider
)

// 在provides中注册session提供者
func Register(name string, provider Provider) {
	if provider == nil {
		log.Fatalf("session:Register %v provider is nil", name)
	}

	if _, dup := provides[name]; dup {
		log.Fatalf("session: Register provider %v is exists", name)
	}

	provides[name] = provider
}

func init() {

	GlobalManager, _ = NewManager("X-Session-ID", "memory", 30*time.Minute)

	go GlobalManager.GC()
}
