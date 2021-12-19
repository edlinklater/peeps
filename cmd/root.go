package cmd

import (
	"errors"
	"fmt"
	"github.com/edlinklater/peeps/model"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var dbConn *gorm.DB = nil

var rootCmd = &cobra.Command{
	Use:   "peeps",
	Short: "keep track of the people you know",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func connectedDb() *gorm.DB {
	if dbConn == nil {
		var connectErr error
		configDir, _ := os.UserConfigDir()
		appData := filepath.Join(configDir, "peeps")

		if _, err := os.Stat(appData); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(appData, 0700)
			if err != nil {
				panic(fmt.Sprintf("Couldn't create data directory %s", appData))
			}
		}

		dbConn, connectErr = gorm.Open(sqlite.Open(filepath.Join(appData, "peeps.db")), &gorm.Config{})
		if connectErr != nil {
			panic("failed to connect database")
		}

		migrateErr := dbConn.AutoMigrate(&model.Peep{}, &model.Nickname{}, &model.Note{})
		if migrateErr != nil {
			panic("failed to set up database")
		}
	}

	return dbConn
}

func strToUint(str string) uint {
	re := regexp.MustCompile(`\D*`)
	stripped := re.ReplaceAllString(str, "")
	integer, _ := strconv.Atoi(stripped)
	return uint(integer)
}
