package tgbot

import (
	"github.com/aliforever/go-genericmap"
	tgbotapi "github.com/aliforever/go-telegram-bot-api"
	"sync"
)

type TgBot struct {
	m      sync.Mutex
	token  string
	menus  *genericmap.GenericMap[menu]
	client *tgbotapi.TelegramBot
}

func New(token string) (tgbot *TgBot, err error) {
	var client *tgbotapi.TelegramBot
	client, err = tgbotapi.NewTelegramBot(token)
	if err != nil {
		return
	}
	tgbot = &TgBot{token: token, menus: genericmap.New[menu](), client: client}
	return
}

func (t *TgBot) AddMenu(menu *menu) {
	t.m.Lock()
	defer t.m.Unlock()

	t.menus.SetPointer(menu.name, menu)
}
