package handler

import (
	"encoding/base64"
	"strings"
)

func base64Decoder(str string) (string, error) {
	fmtstr := strings.Trim(str, " ")
	data, err := base64.StdEncoding.DecodeString(fmtstr)
	if !(fmtstr == "") {
		return "`" + string(data) + "`", err
	} else {
		return string(data), err
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
