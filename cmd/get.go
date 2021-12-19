package cmd

import (
	"fmt"
	"github.com/edlinklater/peeps/model"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "view information for specified peep",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db := connectedDb()

		var peep model.Peep
		peepId := strToUint(args[0])

		res := db.First(&peep, peepId)
		if res.Error != nil {
			log.Fatal(fmt.Sprintf("Couldn't fetch Peep #%d", peepId))
		}

		fmt.Printf("ID:         %d\n", peep.ID)
		fmt.Printf("Real name:  %s\n", peep.RealName)
	},
}
