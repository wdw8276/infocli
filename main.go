package main

// Time    :   2024-06-04 02:02:00 PM

import (
	"fmt"
	"infocli/utils"
	"os"
	"os/user"

	"gorm.io/gorm"
)
const ()

var (
	logger = utils.Logger // logger instance

	// global flags
	gDbFile      string
	gID          int
	gKey         string
	gShowDetail  bool
	gShowVersion bool
	gDebug       bool

	// global variables
	gDb  *gorm.DB
	gErr error
	gRet *gorm.DB
)

func PrintGlobalFlags() {
	if !gDebug {
		return
	} else {
		logger.Println("debug global flags:")
		logger.Println("  db file:", gDbFile)
		logger.Println("  id:", gID)
		logger.Println("  show detail:", gShowDetail)
		logger.Println("  show version:", gShowVersion)
		logger.Println("  debug mode:", gDebug)
	}
}

func main() {
	// get login user name
	user, err := user.Current()
	if err != nil {
		logger.Fatalln(err)
	}

	// default db file: like ~/.fish.db
	defaultDb := user.HomeDir + "/." + user.Username + ".db"

	// add global flags
	rootCmd.PersistentFlags().StringVarP(&gDbFile, "file", "f", defaultDb, "database file")
	rootCmd.PersistentFlags().IntVarP(&gID, "id", "i", 0, "input record id")
	rootCmd.PersistentFlags().BoolVarP(&gShowDetail, "detail", "d", false, "show detail data")
	rootCmd.PersistentFlags().BoolVarP(&gShowVersion, "version", "v", false, "show version")
	rootCmd.PersistentFlags().BoolVarP(&gDebug, "debug", "D", false, "enable debug mode")
	rootCmd.PersistentFlags().StringVarP(&gKey, "key", "k", os.Getenv("INFOCLI_KEY"), "encryption key (or set INFOCLI_KEY env)")

	// add sub commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initDbCmd)
	rootCmd.AddCommand(addCmd)

	queryCmd.AddCommand(nameCmd)
	queryCmd.AddCommand(dataCmd)
	queryCmd.AddCommand(idCmd)
	rootCmd.AddCommand(queryCmd)

	updateCmd.AddCommand(updateNameCmd)
	updateCmd.AddCommand(updateDataCmd)
	rootCmd.AddCommand(updateCmd)

	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(countCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	PrintGlobalFlags()
	os.Exit(0)
}
