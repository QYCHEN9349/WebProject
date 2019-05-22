package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"sync"
	"time"
)

type Manager struct {
	sessionMap  map[string]Session
	lock        sync.Mutex
	maxLifeTime time.Duration
}

func NewManager(maxLifeTime time.Duration) (*Manager, error) {
	return &Manager{maxLifeTime: maxLifeTime}, nil
}
func GenerateSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (p *Manager) SetMaxLiftTime(maxLiftTime time.Duration) {
	p.maxLifeTime = maxLiftTime;
}
func (p *Manager) createSession() {
	sessionId := GenerateSessionId()
	session := Session{id: sessionId, timeAccessed: time.Now()}
	p.lock.Lock()
	p.sessionMap[sessionId] = session
	p.lock.Unlock()
}

func (p *Manager) deleteSession(sessionId string) {
	p.lock.Lock()
	delete(p.sessionMap, sessionId)
	p.lock.Unlock()
}

func (p *Manager) GC() {
	p.lock.Lock()
	for k, v := range p.sessionMap {
		if p.maxLifeTime > time.Now().Sub(v.timeAccessed) {
			delete(p.sessionMap, k)
		}
	}
	p.lock.Unlock()
}
