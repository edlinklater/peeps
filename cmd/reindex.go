package cmd

import (
	"fmt"
	"github.com/blevesearch/bleve/v2"
	"github.com/edlinklater/peeps/model"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func init() {
	rootCmd.AddCommand(reindexCmd)
}

var reindexCmd = &cobra.Command{
	Use:   "reindex",
	Short: "force full rebuild of the search index",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		db := connectedDb()

		reindexResult := reindex(db)
		if reindexResult != nil {
			docCount, _ := reindexResult.DocCount()
			fmt.Printf("%d peeps reindexed.\n", docCount)
		} else {
			log.Fatal("Reindexing failed")
		}
	},
}

func reindex(db *gorm.DB) bleve.Index {
	idx := getIndex()

	var peeps []model.Peep
	dbResult := db.Find(&peeps)
	if dbResult.Error != nil {
		log.Fatal(fmt.Sprintf("Couldn't load Peeps: %s", dbResult.Error))
	}

	for _, peep := range peeps {
		err := idx.Index(strconv.Itoa(int(peep.ID)), peep)
		if err != nil {
			log.Fatal(fmt.Sprintf("Couldn't index Peep: %s", err))
		}
	}

	return idx
}
