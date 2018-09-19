package common

import (
	"sync"

	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// Session is a simple session
type Session struct {
	m *sync.Map
}

// SetUser set user info to session
func (s *Session) SetUser(key interface{}, user models.CommonUser) {
	s.m.Store(key, user)
}

// GetUser get user from session
func (s *Session) GetUser(key interface{}) (models.CommonUser, bool) {
	if v, ok := s.m.Load(key); ok {
		if user, ok := v.(models.CommonUser); ok {
			return user, true
		}
	}
	return models.CommonUser{}, false
}

// Delete delete user from session
func (s *Session) Delete(key interface{}) {
	s.m.Delete(key)
}

var (
	// GlobalSession global session
	GlobalSession *Session
)

func init() {
	GlobalSession = &Session{
		m: new(sync.Map),
	}
}
