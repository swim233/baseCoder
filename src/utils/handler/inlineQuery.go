package handler

import (
	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	"github.com/swim233/baseCoder/utils"
)

func InlineQueryHandler(update tgbotapi.Update) error {
	input := update.InlineQuery.Query

	encodingData := base64Encoder(input)
	encodingArticle := tgbotapi.NewInlineQueryResultArticleMarkdownV2(
		"1", // 唯一ID
		"编码结果",
		"Base64编码结果\n"+encodingData,
	)
	encodingArticle.Description = "base64编码结果预览" + encodingData

	decodingData, err, successTimes := base64Decoder(input)
	decodingArticle := tgbotapi.NewInlineQueryResultArticleHTML(
		"2", // 唯一ID
		"解码结果",
		"Base64解码结果\n"+decodingData,
	)
	decodingArticle.Description = "base64解码结果预览" + decodingData

	results := []interface{}{encodingArticle}
	if err == nil || successTimes != 0 {
		results = append(results, decodingArticle)
	}
	// 构造响应
	inlineConfig := tgbotapi.InlineConfig{
		InlineQueryID: update.InlineQuery.ID,
		Results:       results,
		CacheTime:     0, // 缓存时间（秒）
	}
	utils.Bot.Send(inlineConfig)
	// 发送响应
	return nil
}
