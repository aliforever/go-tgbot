package tgbot

type StateStorage interface {
	GetUserState(userID int64)
	StoreUserState(userID int64, state string)
}
