package cmd

import (
	"fmt"
	"github.com/edlinklater/peeps/model"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "add or edit detail of a peep",
	Run: func(cmd *cobra.Command, args []string) {
		db := connectedDb()

		var peep model.Peep
		peepId := strToUint(args[0])

		res := db.First(&peep, peepId)
		if res.Error != nil {
			log.Fatal(fmt.Sprintf("Couldn't fetch Peep #%d", peepId))
		}

		prompt := promptui.Select{
			Label: fmt.Sprintf("Editing #%d", peep.ID),
			Items: []string{
				fmt.Sprintf("Real name: %s", peep.RealName),
				"❌ Cancel",
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("You choose %q\n", result)
	},
}
