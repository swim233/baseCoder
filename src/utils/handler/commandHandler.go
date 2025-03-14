package handler

import (
	"os"
	"strconv"

	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	"github.com/swim233/baseCoder/utils"
	"github.com/swim233/baseCoder/utils/network"
)

func EncodeCommand(update tgbotapi.Update) error {
	input := update.Message
	if input.ReplyToMessage != nil && input.ReplyToMessage.Text != "" {
		ReplyEncodeCommand(update)
		return nil
	} else {
		if (update.Message.CommandArguments() == "") && (network.GetFileID(update) == "") {
			sendEncodingData(update, "请输入需要编码的内容或回复一条消息")
		} else if (update.Message.CommandArguments() == "") && (network.GetFileID(update) != "") {
			ReplyFileEncodeCommand(update)
		}
		data := base64Encoder(input.CommandArguments())
		sendEncodingData(update, data)
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
			if update.Message.CommandArguments() == "" {
				sendEncodingData(update, "请输入需要编码的内容或回复一条消息")
			}
			sendEncodingData(update, "Base64解码结果\n"+data)
			return nil
		}
		sendEncodingData(update, "格式有误")

	}
	return nil
}

func ReplyEncodeCommand(update tgbotapi.Update) error {
	input := update.Message.ReplyToMessage
	if input != nil {
		data := base64Encoder(input.Text)
		sendEncodingData(update, "Base64编码结果\n"+data)
	}
	return nil

}

func ReplyDecodeCommand(update tgbotapi.Update) error {
	input := update.Message
	data, err := base64Decoder(input.ReplyToMessage.Text)
	if err == nil {
		sendEncodingData(update, "Base64解码结果\n"+data)
	} else {
		sendEncodingData(update, "格式有误")
	}
	return nil
}

// 发送编码文件
func ReplyFileEncodeCommand(update tgbotapi.Update) error {
	data, err := FileEncoder(update)
	if err != nil {
		sendEncodingData(update, err.Error())
		return err
	}
	FileSender(data, update)
	return nil
}

// 发送编码结果
func sendEncodingData(update tgbotapi.Update, data string) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, data)
	msg.ParseMode = tgbotapi.ModeMarkdownV2
	msg.AllowSendingWithoutReply = true
	msg.ReplyToMessageID = int(update.Message.MessageID)
	utils.Bot.Send(msg)
	return nil
}

// 编码文件
func FileEncoder(update tgbotapi.Update) (string, error) {
	file, err := network.DownloadFile(update)
	if err != nil {
		return "", err
	}
	data := base64FileEncoder(file)
	return data, nil
}

// 发送文件
func FileSender(base64 string, update tgbotapi.Update) {
	fileName := strconv.Itoa(update.Message.MessageID) + ".txt"
	context := base64
	os.WriteFile(fileName, []byte(context), 0644)

	msg := tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FilePath(fileName))
	msg.AllowSendingWithoutReply = true
	msg.ReplyToMessageID = int(update.Message.MessageID)
	utils.Bot.Send(msg)

	os.Remove(fileName)
}
