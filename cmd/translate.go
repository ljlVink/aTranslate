/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"aTranslate/conf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
func readConfig() {
	err:=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
	var config  conf.Yaml_config
	viper.Unmarshal(&config)
	fmt.Println(config.General.Openai_key)	
	fmt.Println(config.General.Openai_url)
}
// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "translate",
	Long: `Translate pdf to Chinese.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("translate called")
		readConfig()
		
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)
}
