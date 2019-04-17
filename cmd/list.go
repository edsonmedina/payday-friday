package cmd

import (
	"fmt"
	"time"

	"github.com/edsonmedina/payday-friday/internal/payday"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list next paydays",
	Run: func(cmd *cobra.Command, args []string) {

		now := time.Now()
		nextPaydayFriday := payday.LastFridayOf(now.Month(), now.Year())

		daysLeft := nextPaydayFriday - now.Day()
		weeks := float32(daysLeft) / 7

		fmt.Printf(
			"%s %d:\tpayday is %02d (%.1f weeks / %d days left)\n",
			now.Month().String(),
			now.Year(),
			nextPaydayFriday,
			weeks,
			daysLeft,
		)

		for n := 1; n < 11; n++ {
			m := time.Month(int(now.Month()) + n)
			upcomingPayday := payday.LastFridayOf(m, now.Year())
			paydayTime := time.Date(now.Year(), m, upcomingPayday, 0, 0, 0, 0, time.UTC)

			weeksLong := payday.WeeksLong(m, now.Year())

			fmt.Printf(
				"%s %d:\tpayday is %02d (month is %d weeks long)\n",
				paydayTime.Month().String(),
				paydayTime.Year(),
				paydayTime.Day(),
				weeksLong,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
