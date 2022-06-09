package main

import (
	"fmt"
	tgbotapi "github.com/aliforever/go-telegram-bot-api"
	"github.com/aliforever/go-tgbot"
)

func main() {
	token := "5416326964:AAFLDSnMsOFv8g7yxqnJNHqkygqqMccEfZQ"
	bot, err := tgbot.New(token, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.AddMenu(tgbot.NewMenu("Welcome", func(update *tgbotapi.Update, isSwitched bool) (nextMenu string) {
		text := update.PrivateMessageText()
		if !isSwitched && text != "" {
			if text == "Switch Menu" {
				return "Switch"
			}
		}

		keyboard := bot.API().Tools.Keyboards.NewReplyKeyboardFromSlicesOfStrings([][]string{
			{"Switch Menu"},
		})

		bot.API().Send(bot.API().Message().SetChatId(update.From().Id).SetText("You're in Welcome Menu").SetReplyMarkup(keyboard))
		return
	}))

	bot.AddMenu(tgbot.NewMenu("Switch", func(update *tgbotapi.Update, isSwitched bool) (nextMenu string) {
		if !isSwitched {
			if update.PrivateMessageText() == "Back" {
				return "Welcome"
			}
		}

		bot.API().Send(bot.API().Message().SetChatId(update.From().Id).SetText("You're in Switched Menu").SetReplyMarkup(bot.BackReplyMarkupKeyboard()))
		return
	}))

	bot.GetUpdates()
	fmt.Println(bot)
}
