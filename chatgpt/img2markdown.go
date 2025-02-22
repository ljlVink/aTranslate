package chatgpt

import (
	"aTranslate/conf"
	"os"

	openai "github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var oaikey string
var oaiurl string
var oaimodel string

func loadConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	var config conf.Yaml_config
	viper.Unmarshal(&config)
	oaikey = config.General.Openai_key
	oaiurl = config.General.Openai_url
	oaimodel = config.General.Openai_model
}

func Img2MarkdownOut(base64EncodedString string, OutPutPath string, CurrentPage int) error {
	loadConfig()
	log.Println("calling Img2MarkdownOut -> ImgOutPutPath:", OutPutPath)
	Oaiclient := NewOpenAIClient(oaiurl, oaikey)
	MultiContent := []openai.ChatMessagePart{
		{
			Type: openai.ChatMessagePartTypeText,
			Text: conf.Prompt,
		},
		{
			Type: openai.ChatMessagePartTypeImageURL,
			ImageURL: &openai.ChatMessageImageURL{
				URL:    base64EncodedString,
				Detail: openai.ImageURLDetailAuto,
			},
		},
	}
	resp, err := Oaiclient.SendMessage(MultiContent, oaimodel)
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Println("Page:", CurrentPage, "Translated Successfully")
	}

	if err := os.WriteFile(OutPutPath, []byte(resp), 0666); err != nil {
		log.Fatalln(err)
	}
	return nil
}
