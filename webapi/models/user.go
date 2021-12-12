package models

import (
	"sync"
	"time"

	"github.com/furusax0621/goa-sample/webapi/app"
)

type SyncUserMap struct {
	mx     sync.RWMutex
	nextID int
	users  map[int]*app.UserMedia
}

func NewSyncUserMap() *SyncUserMap {
	return &SyncUserMap{
		nextID: 1,
		users:  make(map[int]*app.UserMedia),
	}
}

func (m *SyncUserMap) Store(givenName, familyName string, createdAt time.Time) {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.users[m.nextID] = &app.UserMedia{
		UserID:     m.nextID,
		GivenName:  givenName,
		FamilyName: familyName,
		CreatedAt:  createdAt,
	}
	m.nextID++
}

func (m *SyncUserMap) Load(id int) (*app.UserMedia, bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()

	user, ok := m.users[id]
	return user, ok
}
