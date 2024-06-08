package cmd

import (
	"fmt"
	"net/http"

	utils "vlr-cli/helper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var live bool

var scoreCmd = &cobra.Command{
	Use:     "score",
	Short:   "Get the latest scores of live Valorant matches.",
	Long:    `This command allows you to get the latest scores of Live and Upcoming Valorant matches.`,
	Example: `vlr score`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Wait for the results to load...")
		getScoresFromApi()
	},
}

func init() {
	scoreCmd.Flags().BoolP("live", "l", false, "Get the scores of all Live matches.")
	viper.BindPFlag("live", scoreCmd.Flags().Lookup("live"))
	rootCmd.AddCommand(scoreCmd)
}

func getScoresFromApi() {
	res, err := http.Get("http://vlr-api.centralindia.cloudapp.azure.com/matches")
	if err != nil {
		fmt.Printf("Error fetching scores: %v\n", err)
		return
	}
	defer res.Body.Close()

	live = viper.GetBool("live")
	count := viper.GetInt("count")
	region := viper.GetString("region")
	utils.ParseResponse(res.Body, live, count, region)
}
