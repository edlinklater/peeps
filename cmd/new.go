package cmd

import (
	"fmt"
	"github.com/Bowery/prompt"
	"github.com/edlinklater/peeps/model"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "add a new peep",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		db := connectedDb()

		realName, _ := prompt.Basic("Real name?", true)

		peep := model.Peep{RealName: realName}
		peepResult := db.Create(&peep)

		if peepResult.Error != nil {
			log.Fatal(fmt.Sprintf("Couldn't create Peep: %s", peepResult.Error))
		}

		fmt.Printf("Peep created: #%d", peep.ID)

		idx := getIndex()
		_ = idx.Index(cast.ToString(peep.ID), peep)
	},
}
