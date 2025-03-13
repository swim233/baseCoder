package handler

import (
	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	"github.com/swim233/baseCoder/utils"
)

func EncodeCommand(update tgbotapi.Update) error {
	input := update.Message
	if input.ReplyToMessage != nil {
		ReplyEncodeCommand(update)
		return nil
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Base64编码结果\n"+base64Encoder(input.CommandArguments()))
		msg.ParseMode = tgbotapi.ModeMarkdownV2
		msg.ReplyToMessageID = update.Message.MessageID
		msg.AllowSendingWithoutReply = true
		utils.Bot.Send(msg)
		return nil
	}
}
func DecodeCommand(update tgbotapi.Update) error {
	input := update.Message
	if input.ReplyToMessage != nil {
		ReplyDecodeCommand(update)
		return nil
	} else {
		data, err := base64Decoder(input.CommandArguments())
		if err == nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Base64解码结果\n"+data)
			msg.ParseMode = tgbotapi.ModeMarkdownV2
			msg.ReplyToMessageID = update.Message.MessageID
			msg.AllowSendingWithoutReply = true
			utils.Bot.Send(msg)
			return nil
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "格式有误")
			utils.Bot.Send(msg)
			return nil
		}
	}
}

func ReplyEncodeCommand(update tgbotapi.Update) error {
	input := update.Message.ReplyToMessage
	if input != nil {
		data := base64Encoder(input.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Base64编码结果\n"+data)
		msg.ParseMode = tgbotapi.ModeMarkdownV2
		msg.AllowSendingWithoutReply = true
		msg.ReplyToMessageID = int(update.Message.MessageID)
		utils.Bot.Send(msg)
		return nil
	} else {
		return nil
	}
}
func ReplyDecodeCommand(update tgbotapi.Update) error {
	input := update.Message
	data, err := base64Decoder(input.ReplyToMessage.Text)
	if err == nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Base64解码结果\n"+data)
		msg.ParseMode = tgbotapi.ModeMarkdownV2
		msg.ReplyToMessageID = update.Message.MessageID
		msg.AllowSendingWithoutReply = true
		utils.Bot.Send(msg)
		return nil
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "格式有误")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.AllowSendingWithoutReply = true
		utils.Bot.Send(msg)
		return nil
	}

}
