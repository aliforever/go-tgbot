package tgbot

import "sync"

type StateStorage interface {
	GetUserState(userID int64) (state string)
	StoreUserState(userID int64, state string)
}

type temporaryStateStorage struct {
	sync.Map
}

func newTemporaryStateStorage() *temporaryStateStorage {
	return &temporaryStateStorage{}
}

func (t *temporaryStateStorage) GetUserState(userID int64) (state string) {
	if val, ok := t.Load(userID); ok {
		if valStr, ok := val.(string); ok {
			return valStr
		}
	}

	return "Welcome"
}

func (t *temporaryStateStorage) StoreUserState(userID int64, state string) {
	t.Store(userID, state)
}
