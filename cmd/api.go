package cmd

import (
	"log"

	"ebook/pkg/api"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Root short description",
	Long:  "Root long description",
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Root short description",
	Long:  "Root long description",
	Run:   startAPI,
}

func startAPI(*cobra.Command, []string) {
	// fmt.Println("hellooo...")
	api.Start()

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
