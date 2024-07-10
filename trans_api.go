package main

import (
	"encoding/json"
	"github.com/yrshiben/deepl-translation-workflow/model"
	"io"
	"net/http"
	"strings"
)

const TransApiUrl string = "https://api-free.deepl.com/v2/translate"

type TransApi struct {
	AuthKey string
}

type TranslateRequestParams struct {
	Text       []string `json:"text"`
	SourceLang string   `json:"source_lang"`
	TargetLang string   `json:"target_lang"`
}

func NewTransApi(authKey string) *TransApi {
	return &TransApi{AuthKey: authKey}
}

// GetTransResult from/to: EN or ZH
func (api *TransApi) GetTransResult(query, from, to string) (*model.Result, error) {
	params := TranslateRequestParams{Text: []string{query}, SourceLang: from, TargetLang: to}
	bytes, _ := json.Marshal(params)
	req, err := http.NewRequest("GET", TransApiUrl, strings.NewReader(string(bytes)))
	var queryResult *model.Result
	if err != nil {
		return queryResult, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "DeepL-Auth-Key "+api.AuthKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return queryResult, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if result, err := io.ReadAll(resp.Body); err != nil {
		return queryResult, err
	} else {
		err = json.Unmarshal(result, &queryResult)
		return queryResult, err
	}
}
