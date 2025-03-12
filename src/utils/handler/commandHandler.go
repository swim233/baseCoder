package handler

import (
	"github.com/ijnkawakaze/telegram-bot-api"
	"github.com/swim233/baseCoder/utils"
)

func EncodeCommand(update tgbotapi.Update) error {
	input := update.Message
	msg := tgbotapi.NewMessage(update.Message.From.ID, "Base64编码结果\n"+base64Encoder(input.CommandArguments()))
	msg.ParseMode = tgbotapi.ModeMarkdownV2
	utils.Bot.Send(msg)
	return nil
}
func DecodeCommand(update tgbotapi.Update) error {
	input := update.Message
	data, err := base64Decoder(input.CommandArguments())
	if err == nil {
		msg := tgbotapi.NewMessage(update.Message.From.ID, "Base64解码结果\n"+data)
		msg.ParseMode = tgbotapi.ModeMarkdownV2
		utils.Bot.Send(msg)
		return nil
	} else {
		msg := tgbotapi.NewMessage(update.Message.From.ID, "格式有误")
		utils.Bot.Send(msg)
		return nil
	}
}

//func ReplyEncodeCommand(update tgbotapi.Update) error{

//}
