package tgbot

import (
	tgbotapi "github.com/aliforever/go-telegram-bot-api"
	"github.com/aliforever/go-telegram-bot-api/structs"
	"sync"
)

type TgBot struct {
	m               sync.Mutex
	token           string
	menus           sync.Map
	client          *tgbotapi.TelegramBot
	stateStorage    StateStorage
	defaultResponse string
}

func New(token string, stateStorage StateStorage) (tgbot *TgBot, err error) {
	var client *tgbotapi.TelegramBot
	client, err = tgbotapi.NewTelegramBot(token)
	if err != nil {
		return
	}
	if stateStorage == nil {
		stateStorage = newTemporaryStateStorage()
	}
	tgbot = &TgBot{token: token, menus: sync.Map{}, client: client, stateStorage: stateStorage, defaultResponse: "Command not found!"}
	return
}

func (t *TgBot) SetDefaultResponse(response string) {
	t.defaultResponse = response
}

func (t *TgBot) AddMenu(name string, handler menuHandler) {
	t.m.Lock()
	defer t.m.Unlock()

	t.menus.Store(name, handler)
}

func (t *TgBot) BackReplyMarkupKeyboard() *structs.ReplyKeyboardMarkup {
	return t.client.Tools.Keyboards.NewReplyKeyboardFromSlicesOfStrings([][]string{
		{
			"Back",
		},
	}).SetResizeKeyboard(true)
}

func (t *TgBot) API() (api *tgbotapi.TelegramBot) {
	return t.client
}

func (t *TgBot) getMenu(name string) (m menuHandler, exists bool) {
	t.m.Lock()
	defer t.m.Unlock()

	var val any
	val, exists = t.menus.Load(name)
	if exists {
		m = val.(menuHandler)
	}

	return
}

func (t *TgBot) GetUpdates() {
	for update := range t.client.GetUpdates().LongPoll() {
		text := update.PrivateMessageText()
		if text != "" {
			userState := t.stateStorage.GetUserState(update.Message.From.Id)
			if m, ok := t.getMenu(userState); ok {
				if nextMenu := m(&update, false); nextMenu != "" {
					if m, ok := t.getMenu(nextMenu); ok {
						if err := t.stateStorage.StoreUserState(update.Message.From.Id, nextMenu); err == nil {
							m(&update, true)
							continue
						} else {
							t.client.Send(t.client.Message().SetChatId(update.Message.From.Id).SetText("Error updating user state"))
						}
					}
					t.client.Send(t.client.Message().SetChatId(update.Message.From.Id).SetText(t.defaultResponse))
				}
				continue
			}
			t.client.Send(t.client.Message().SetChatId(update.Message.From.Id).SetText(t.defaultResponse))
		}
	}
}
