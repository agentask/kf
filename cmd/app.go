/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}
		
		createApp(app)
	},
}

func init() {
	appCmd.PersistentFlags().String("name", "", "the app name")
	rootCmd.AddCommand(appCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createApp(name string) {
	params := map[string]string{
		"AppName": name,
	}
	os.MkdirAll(filepath.Join(name, "config"), os.ModePerm)

	// Create main.go, go.mod, and config/config.go using embedded templates
	generateFileFromTemplate(filepath.Join(name, "main.go"), "templates/main.go.tmpl", params)
	generateFileFromTemplate(filepath.Join(name, "config/config.go"), "templates/config.go.tmpl", params)

	fmt.Printf("App '%s' created with main.go and config/config.go.\n", name)
}
