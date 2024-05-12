package cmd

import (
	"fmt"
	"net/http"

	utils "VLR-CLI/helper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var live bool

var scoreCmd = &cobra.Command{
	Use:     "score",
	Short:   "Get the latest scores of Valorant matches.",
	Long:    `This command allows you to get the latest scores of Live and Upcoming Valorant matches.`,
	Example: `vlr score`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Wait for the scores to load...")
		res, err := http.Get("http://localhost:8080/matches")
		if err != nil {
			fmt.Printf("Error fetching scores: %v\n", err)
			return
		}
		defer res.Body.Close()

		// fmt.Println("----------------------- Scores -----------------------")
		live = viper.GetBool("live")
		count := viper.GetInt("count")
		region := viper.GetString("region")
		utils.ParseResponse(res.Body, live, count, region)
	},
}

func init() {
	scoreCmd.Flags().BoolP("live", "l", false, "Get the scores of all Live matches.")
	scoreCmd.Flags().Int16P("count", "c", 5, "Number of matches to display.")
	scoreCmd.Flags().StringP("region", "r", "", "Filter matches by region. Available regions: [eu, na, apac, cn]")
	viper.BindPFlag("live", scoreCmd.Flags().Lookup("live"))
	viper.BindPFlag("count", scoreCmd.Flags().Lookup("count"))
	viper.BindPFlag("region", scoreCmd.Flags().Lookup("region"))
	rootCmd.AddCommand(scoreCmd)
}
