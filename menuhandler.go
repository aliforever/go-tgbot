package tgbot

import tgbotapi "github.com/aliforever/go-telegram-bot-api"

// menuHandler
type menuHandler func(update *tgbotapi.Update, isSwitched bool) (nextMenu string)
