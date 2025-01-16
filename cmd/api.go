package cmd

import (
	"log"

	"ebook/app/database"
	"ebook/pkg/api"

	"ebook/app"

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
	
	gormDB, _, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("couldn't connect database..")
	}


	r := app.APIRouter(gormDB)
	api.Start(r)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
