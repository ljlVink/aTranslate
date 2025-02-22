/*
Copyright © 2025 ljlvink
*/
package cmd

import (
	"aTranslate/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// Root 命令
var rootCmd = &cobra.Command{
	Use:   "aTranslate",
	Short: "A tool to translate PDFs using the OpenAI platform.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,                        // 显示完整时间戳
		TimestampFormat: "2006-01-02 15:04:05-07:00", // 自定义时间戳格式
		ForceColors:     true,                        // 强制启用颜色
		DisableColors:   false,                       // 不禁用颜色
	})
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is aTranslate.yaml)")
}

// 初始化配置
func initConfig() {
	workingDir := "." // 默认为当前目录
	if cfgFile == "" {
		viper.AddConfigPath(workingDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("aTranslate")
	} else {
		viper.SetConfigFile(cfgFile)
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else if cfgFile != "" {
		log.Fatalln("Failed to read config file:", err)
	}
	// Check if workingDir/outputs exists, if not, create it
	outputsDir := fmt.Sprintf("%s/outputs", workingDir)
	if !utils.IsDirExist(outputsDir) {
		if err := os.Mkdir(outputsDir, os.ModePerm); err != nil {
			log.Fatalln("Failed to create outputs directory:", err)
		}
	}
}
