package tgbot

type menu struct {
	name    string
	handler menuHandler
}

func NewMenu(name string, handler menuHandler) *menu {
	return &menu{
		name:    name,
		handler: handler,
	}
}
