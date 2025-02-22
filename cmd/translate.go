/*
Copyright © 2025 ljlvink
*/
package cmd

import (
	"aTranslate/conf"
	"aTranslate/translate"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config conf.Yaml_config

func readConfig() error {
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.Unmarshal(&config)
	return nil
}

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "translate",
	Long:  `Translate pdf to Chinese.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := readConfig()
		if err != nil {
			log.Fatalln("Error in readConfig ,", err)
		}
		file, _ := cmd.Flags().GetString("file")
		err = translate.DoTranslate(file)
		if err != nil {
			log.Println("Error in DoTranslate ,", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)
	translateCmd.Flags().StringP("file", "f", "", "File to be translated")
	translateCmd.MarkFlagRequired("file")
}
