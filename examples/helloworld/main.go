package main

import (
	"fmt"
	"github.com/aliforever/go-tgbot"
)

func main() {
	token := ""
	bot, err := tgbot.New(token)
	if err != nil {
		return
	}
	fmt.Println(bot)
}
