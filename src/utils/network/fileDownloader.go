package network

import (
	"fmt"
	"io"
	"net/http"

	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	"github.com/swim233/baseCoder/utils"
)

func getUrl(update tgbotapi.Update) (url string, err error) {
	fileID := GetFileID(update)
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
		fileID = update.Message.ReplyToMessage.Document.FileID
	case update.Message != nil && update.Message.ReplyToMessage.Photo != nil:
		fileID = update.Message.ReplyToMessage.Photo[len(update.Message.ReplyToMessage.Photo)-1].FileID
	case update.Message != nil && update.Message.ReplyToMessage.Video != nil:
		fileID = update.Message.ReplyToMessage.Video.FileID
	default:
		fileID = ""
	}
	return fileID
}
