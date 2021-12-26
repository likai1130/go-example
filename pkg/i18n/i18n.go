package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 本地化包
var bundle *i18n.Bundle

func init() {
	// 设置默认语言
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]

	// 加载公共信息语言包文件
	bundle.MustLoadMessageFile(path + "/tomls/nft.en.toml")
	bundle.MustLoadMessageFile(path + "/tomls/nft.zh-cn.toml")
}

/* 本地化语言
param:
	lang 语言 en|zh-cn
	messageID 语言文件中的 messageID
	errCode: string 错误码
	templateData: i18n 文件中需要变量替换的内容
	pluralCount: 传入 int 或 int64 类型数据 ， 根据数字判断是否返回复数格式 msg
*/
func MustLocalize(lang, accept, messageID string, templateData interface{}, pluralCount interface{}) string {
	localizer := i18n.NewLocalizer(bundle, lang, accept)
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
		PluralCount:  pluralCount,
	})
}
