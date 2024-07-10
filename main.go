package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yrshiben/deepl-translation-workflow/model"
	"os"
	"unicode"
)

var ErrItem = model.AlfredItem{Title: "翻译失败", Subtitle: ""}

var api *TransApi

func main() {
	var query, authKey string
	app := &cli.App{
		Name:      "DeepL 翻译工具",
		Usage:     "deepl-translation",
		UsageText: "deepl-translation -q <要翻译的内容> --authKey <authKey>",
		Version:   "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "query", Aliases: []string{"q"}, Usage: "要翻译的内容", Required: true, Destination: &query},
			&cli.StringFlag{Name: "authKey", Usage: "在 https://www.deepl.com/zh/your-account/keys 查看你的 API Key", Required: true, Destination: &authKey},
		},
		Action: func(c *cli.Context) error {
			api = NewTransApi(authKey)
			// 默认 英译汉
			var from, to = "EN", "ZH"
			if IsChinese(query) { // 输入的是汉文，则 汉译英
				from, to = to, from
			}
			var items []model.AlfredItem
			if result, err := api.GetTransResult(query, from, to); err != nil || len(result.Translations) == 0 {
				items = []model.AlfredItem{ErrItem}
			} else {
				data := result.Translations
				items = make([]model.AlfredItem, 0, len(data))
				for _, val := range data {
					items = append(items, model.AlfredItem{
						Title:    val.Text,
						Icon:     model.Icon{Path: "./icon.svg"},
						Subtitle: query,
						Arg:      val.Text,
					})
				}
			}
			alfredList := model.AlfredList{Items: items}
			fmt.Println(alfredList.ToJson())
			return nil
		},
	}
	_ = app.Run(os.Args)
}

// IsChinese 判断字符串中是否包含中文
func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}
