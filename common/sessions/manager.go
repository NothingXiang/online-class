package sessions

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/satori/go.uuid"
)

// sessions 管理器
type Manager struct {
	cookieName string
	sync.RWMutex
	provider    Provider
	maxLifeTime time.Duration
}

// 构造一个新的manager
func NewManager(cookieName string, provider string, maxLifeTime time.Duration) (*Manager, error) {
	provide, ok := provides[provider]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q ", provider)
	}

	return &Manager{cookieName: cookieName, provider: provide, maxLifeTime: maxLifeTime}, nil
}

// 提供一个全局唯一的新的id
func (m *Manager) sessionID() string {
	return uuid.NewV4().String()
}

// 清除过期session
func (m *Manager) GC() {
	m.Lock()
	defer m.Unlock()
	m.provider.SessionGC(m.maxLifeTime)

	// 定时清除过期session
	//time.AfterFunc(m.maxLifeTime, m.GC)
}

// 移除某个session
func (m *Manager) Remove(sid string) error {
	if len(sid) == 0 {
		return nil
	}

	m.Lock()
	defer m.Unlock()

	return m.provider.SessionDestroy(sid)
}

func (m *Manager) SessionStart(r *http.Request) (s Session) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		sid := m.sessionID()
		s, _ = m.provider.SessionInit(sid)

	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		s, _ = m.provider.SessionRead(sid)
	}
	return
}

func (m *Manager) SessionDestroy(r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}

	m.provider.SessionDestroy(cookie.Value)

}
