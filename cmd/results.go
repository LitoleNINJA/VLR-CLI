package cmd

import (
	"fmt"
	"net/http"

	utils "vlr-cli/helper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var resultCmd = &cobra.Command{
	Use:     "result",
	Short:   "Get the latest results of Valorant matches.",
	Long:    `This command allows you to get the latest results of Completed Valorant matches.`,
	Example: `vlr result`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Wait for the results to load...")
		getResultsFromApi()
	},
}

func init() {
	rootCmd.AddCommand(resultCmd)
}

func getResultsFromApi() {
	res, err := http.Get("http://vlr-api.centralindia.cloudapp.azure.com/matches?status=completed")
	if err != nil {
		fmt.Printf("Error fetching scores: %v\n", err)
		return
	}
	defer res.Body.Close()

	count := viper.GetInt("count")
	region := viper.GetString("region")
	utils.ParseResponse(res.Body, false, count, region)
}
