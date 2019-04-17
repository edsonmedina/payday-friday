package cmd

import (
	"fmt"
	"time"

	"github.com/edsonmedina/payday-friday/internal/payday"
	"github.com/spf13/cobra"
)

var whenCmd = &cobra.Command{
	Use:   "when",
	Short: "When is next friday payday?",
	Run: func(cmd *cobra.Command, args []string) {

		now := time.Now()
		nextPaydayFriday := payday.LastFridayOf(now.Month(), now.Year())

		daysLeft := nextPaydayFriday - now.Day()
		weeks := float32(daysLeft) / 7

		fmt.Printf(
			"%s %d: next payday is %02d (%.1f weeks / %d days left)\n",
			now.Month().String(),
			now.Year(),
			nextPaydayFriday,
			weeks,
			daysLeft,
		)
	},
}

func init() {
	rootCmd.AddCommand(whenCmd)
}
