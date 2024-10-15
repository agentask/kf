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

// featureCmd represents the feature command
var featureCmd = &cobra.Command{
	Use:   "feature",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		feature, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}

		createFeature(feature)
	},
}

func init() {
	featureCmd.PersistentFlags().String("name", "", "the feature name")
	rootCmd.AddCommand(featureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// featureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// featureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createFeature(feature string) {
	files := []string{
		"controller.go", "routes.go",
		"service.go", "repository.go","model.go",
	}
	params := map[string]string{
		"FeatureName":            feature,
		"FeatureNameCapitalized": capitalize(feature),
	}

	featurePath := filepath.Join("api", feature)
	os.MkdirAll(filepath.Join("api", feature), os.ModePerm)

	for _, file := range files {
		location := filepath.Join(featurePath, file)
		templateFile := fmt.Sprintf("templates/%s.tmpl", file)
		generateFileFromTemplate(location, templateFile, params)
	}

	fmt.Printf("Feature '%s' created in 'api/%s'.\n", feature, feature)
}
