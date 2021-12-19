package cmd

import (
	"errors"
	"fmt"
	"github.com/blevesearch/bleve/v2"
	"github.com/edlinklater/peeps/model"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "find peeps by name or nickname",
	Run: func(cmd *cobra.Command, args []string) {
		index := getIndex()

		queryString := strings.Join(args, " ")
		query := bleve.NewMatchQuery(queryString)
		search := bleve.NewSearchRequest(query)
		search.SortBy([]string{"_id"})
		searchResults, err := index.Search(search)
		if err != nil {
			log.Fatal(fmt.Sprintf("Search failed: %s", err))
		}

		if searchResults.Total < 1 {
			log.Fatal(fmt.Sprintf("No results found for: %s", queryString))
		}

		var peepIds []uint
		var peeps []model.Peep

		for _, hit := range searchResults.Hits {
			id, _ := strconv.Atoi(hit.ID)
			peepIds = append(peepIds, uint(id))
		}

		db := connectedDb()
		db.Find(&peeps, peepIds)

		for _, peep := range peeps {
			fmt.Printf("#%-6d%s\n", peep.ID, peep.RealName)
		}
	},
}

func getIndex() bleve.Index {
	var idx bleve.Index
	var idxErr error

	cacheDir, _ := os.UserCacheDir()
	searchData := filepath.Join(cacheDir, "peeps")
	_ = os.Mkdir(searchData, 0700)
	indexFile := filepath.Join(searchData, "peeps.bleve")

	if _, err := os.Stat(indexFile); errors.Is(err, os.ErrNotExist) {
		mapping := bleve.NewIndexMapping()
		idx, idxErr = bleve.New(indexFile, mapping)
	} else {
		idx, idxErr = bleve.Open(indexFile)
	}

	if idxErr != nil {
		log.Fatal(fmt.Sprintf("Couldn't open search index: %s", idxErr))
	}

	return idx
}
