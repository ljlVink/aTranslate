/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"aTranslate/chatgpt"
	"aTranslate/conf"
	"aTranslate/utils"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var oaikey string
var oaiurl string

func DDreadConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	var config conf.Yaml_config
	viper.Unmarshal(&config)
	oaikey = config.General.Openai_key
	oaiurl = config.General.Openai_url
	fmt.Println(oaikey)
	fmt.Println(oaiurl)
}

// debugChatCmd represents the debugChat command
var debugChatCmd = &cobra.Command{
	Use:   "debugChat",
	Short: "debugChat",
	Long:  `debug`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("debugChat called")
		DDreadConfig()
		img_base64url, err := utils.Img2Base64Url("/home/sensorfaucet/Desktop/aTranslate/img/123/image-00007.jpg")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(img_base64url)
		Oaiclient := chatgpt.NewOpenAIClient(oaiurl, oaikey)
		MultiContent := []openai.ChatMessagePart{
			{
				Type: openai.ChatMessagePartTypeText,
				Text: conf.Prompt,
			},
			{
				Type: openai.ChatMessagePartTypeImageURL,
				ImageURL: &openai.ChatMessageImageURL{
					URL:    img_base64url,
					Detail: openai.ImageURLDetailAuto,
				},
			},
		}
		resp, err := Oaiclient.SendMessage(MultiContent, openai.GPT4oLatest)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	rootCmd.AddCommand(debugChatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// debugChatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// debugChatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
