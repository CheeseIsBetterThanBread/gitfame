//go:build !solution

package main

import (
	"github.com/spf13/cobra"
	"gitlab.com/slon/shad-go/gitfame/internal/cli"
	"gitlab.com/slon/shad-go/gitfame/internal/stats"
	"log"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gitfame",
		Short: "Gitfame is a utility to calculate statistics of authors in a Git repository",
		Run: func(cmd *cobra.Command, args []string) {
			config := cli.ParseFlags(cmd)
			statistics := stats.CalculateStats(config)
			stats.PrintStats(statistics, config.Format)
		},
	}

	cli.SetupFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("\tCould not calculate statistics: %v", err)
	}
}
