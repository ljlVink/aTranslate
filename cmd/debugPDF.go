/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"aTranslate/pdf2img"
	"fmt"
	"github.com/spf13/cobra"
)

// debugPDFCmd represents the debugPDF command
var debugPDFCmd = &cobra.Command{
	Use:   "debugPDF",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("debugPDF called")
		pdf2img.Pdf2img("./test.pdf")
	},
}

func init() {
	rootCmd.AddCommand(debugPDFCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// debugPDFCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// debugPDFCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
