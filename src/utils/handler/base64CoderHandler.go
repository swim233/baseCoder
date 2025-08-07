package handler

import (
	"encoding/base64"
	"errors"
	"strings"
)

func base64Decoder(str string) (string, error, int) {
	var resultBuilder strings.Builder
	var successTimes int
	formattedString := strings.Fields(str)
	for _, v := range formattedString {
		if decodedText, err := base64.StdEncoding.DecodeString(v); err == nil {
			resultBuilder.Write(decodedText)
			resultBuilder.WriteString("\n")
			successTimes++
		} else {
			resultBuilder.WriteString(v + "\n")
		}
	}
	if successTimes == 0 {
		return "`" + string(str) + "`", errors.New("格式不符合"), 0
	} else {
		return resultBuilder.String(), nil, successTimes
	}

}
func base64Encoder(str string) string {
	if !(str == "") {
		s := []byte(str)
		data := base64.StdEncoding.EncodeToString(s)
		return "`" + data + "`"
	}
	s := []byte(str)
	data := base64.StdEncoding.EncodeToString(s)
	return data

}
func base64FileEncoder(src []byte) string {
	if !(src == nil) {
		data := base64.StdEncoding.EncodeToString(src)
		return data
	}
	return ""
}
