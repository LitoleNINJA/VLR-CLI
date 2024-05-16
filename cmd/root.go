package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "vlr",
	Short:   "Get the latest scores of Valorant matches.",
	Long:    `This command allows you to get the latest scores of Live and Upcoming Valorant matches.`,
	Version: "1.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntP("count", "c", 15, "Number of matches to display.")
	rootCmd.PersistentFlags().StringP("region", "r", "", `Filter matches by region. Available regions: 
	eu (EMEA), na (Americas), 
	apac (Asia Pacific), cn (China), 
	kr (Korea), jp (Japan), 
	br (Brazil), es (Spain), 
	latam (Latin America)`)
	viper.BindPFlag("count", rootCmd.PersistentFlags().Lookup("count"))
	viper.BindPFlag("region", rootCmd.PersistentFlags().Lookup("region"))
}
