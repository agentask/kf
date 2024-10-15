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

// libCmd represents the lib command
var libCmd = &cobra.Command{
	Use:   "lib",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lib, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}

		createLib(lib)
	},
}

func init() {
	libCmd.PersistentFlags().String("name", "", "the library name")
	rootCmd.AddCommand(libCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// libCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// libCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createLib(name string) {
	libPath := filepath.Join("pkg", name)
	params := map[string]string{"LibName": name}
	os.MkdirAll(filepath.Join("pkg", name), os.ModePerm)

	// Create library file using an embedded template
	generateFileFromTemplate(filepath.Join(libPath, fmt.Sprintf("%s.go", name)), "templates/lib.go.tmpl", params)

	fmt.Printf("Library '%s' created in 'pkg/%s'.\n", name, name)
}
