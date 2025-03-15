package network

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	"github.com/swim233/baseCoder/utils"
)

func getUrl(update tgbotapi.Update) (url string, err error) {
	fileID := GetFileID(update)
	if fileID == "无法获取文件" || fileID == "文件大小过大" {
		return "", errors.New(fileID)
	}
	FileURL, err := func(bot tgbotapi.BotAPI, fileID string) (string, error) {
		file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath), nil
	}(*utils.Bot, fileID)
	if err != nil {
		return "", err
	}
	return FileURL, nil
}

func DownloadFile(update tgbotapi.Update) ([]byte, error) {
	url, err := getUrl(update)
	if err != nil {
		return nil, err
	}
	rsps, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsps.Body.Close()
	data, err := io.ReadAll(rsps.Body)
	if err != nil {
		return nil, err
	}
	return data, err
}
func GetFileID(update tgbotapi.Update) string {
	var fileID string
	switch {
	case update.Message != nil && update.Message.ReplyToMessage.Document != nil:
		if update.Message.ReplyToMessage.Document.FileSize > utils.BotConfig.DecodeFileMaxSize {
			return "文件大小过大"
		}
		fileID = update.Message.ReplyToMessage.Document.FileID
	case update.Message != nil && update.Message.ReplyToMessage.Photo != nil:
		if update.Message.ReplyToMessage.Photo[len(update.Message.ReplyToMessage.Photo)-1].FileSize > utils.BotConfig.DecodeFileMaxSize {
			return "文件大小过大"
		}
		fileID = update.Message.ReplyToMessage.Photo[len(update.Message.ReplyToMessage.Photo)-1].FileID
	case update.Message != nil && update.Message.ReplyToMessage.Video != nil:
		if update.Message.ReplyToMessage.Video.FileSize > utils.BotConfig.DecodeFileMaxSize {
			return "文件大小过大"
		}
		fileID = update.Message.ReplyToMessage.Video.FileID
	default:
		fileID = "无法获取文件 "
	}
	return fileID
}
