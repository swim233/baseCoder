package main

import (
	"github.com/swim233/baseCoder/utils"
	"github.com/swim233/baseCoder/utils/handler"
)

func main() {
	utils.InitBot()
	utils.Bot.Debug = true
	b := utils.Bot.AddHandle()
	b.NewInlineQueryProcessor("", handler.InlineQueryHandler)
	b.Run()
}
