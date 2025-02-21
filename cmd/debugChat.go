/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"aTranslate/chatgpt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"aTranslate/conf"
)
var oaikey string
var oaiurl string

func DDreadConfig() {
	err:=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
	var config  conf.Yaml_config
	viper.Unmarshal(&config)
	oaikey=config.General.Openai_key
	oaiurl=config.General.Openai_url
	fmt.Println(oaikey)
	fmt.Println(oaiurl)
}



// debugChatCmd represents the debugChat command
var debugChatCmd = &cobra.Command{
	Use:   "debugchat",
	Short: "debugchat",
	Long: `debug`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("debugChat called")
		DDreadConfig()
		oaiclient:=chatgpt.NewOpenAIClient(oaiurl,oaikey)
		messages:=[]chatgpt.Message{
			{
				Role:    "system",
				Content: "你是一个翻译机器，你需要翻译这个图片成中文意义",
			},
			{
				Role:    "user",
				Content: "翻译中文",
			},
		}
		answer,err:=oaiclient.SendMessage(messages,"chatgpt-4o-latest",`C:\Users\16662\Desktop\atranslate\image.png`)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(answer)
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
