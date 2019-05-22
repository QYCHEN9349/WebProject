package session

import "time"

type Session struct {
	id           string
	timeAccessed time.Time
	value        map[string]string
}

func (p *Session) UpdateTimeAccessed() {
	p.timeAccessed = time.Now()
}
func (p *Session) SetKV(key string, value string) {
	p.value[key] = value;
}
