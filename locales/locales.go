package locales

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// init
func init() {
	initEn(language.Make("en"))
	initZh(language.Make("zh"))
}

// initEn will init en support.
func initEn(tag language.Tag) {
	message.SetString(tag, "0", "Success")
	message.SetString(tag, "EX001", "Fail")
}

// initZh will init zh support.
func initZh(tag language.Tag) {
	message.SetString(tag, "0", "操作成功")
	message.SetString(tag, "EX001", "Fail")
}
