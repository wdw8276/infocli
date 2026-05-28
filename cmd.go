package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func initConfig() {
	if gShowVersion {
		logger.Println(gName, gVersion)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

var rootCmd = &cobra.Command{
	Use:   "infocli",
	Short: "A simple tool to store and query data, based on sqlite3.\nAuthor: diwen @Copyright 2024",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("Welcome to infocli! -h for help.")
	},
}

var versionCmd = &cobra.Command{
	Use:   "v",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println(gName, gVersion)
	},
}

var initDbCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database",
	Run: func(cmd *cobra.Command, args []string) {
		InitDB()
	},
}

var addCmd = &cobra.Command{
	Use:   "a",
	Short: "Add a record, first as name, second as data, or read from pipeline",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		data := ""
		if len(args) > 1 {
			data = args[1]
		} else {
			// read data from stdin pipeline
			reader := bufio.NewReader(os.Stdin)
			inputBytes, err := io.ReadAll(reader)
			if err != nil {
				logger.Fatalln("read pipeline data failed:", err)
			} else {
				data = string(inputBytes)
			}
		}
		logger.Println("Adding record:", name, data)
		SaveInfo(name, data)
	},
}

var queryCmd = &cobra.Command{
	Use:   "q",
	Short: "Query data by input name, or using subcommands to query by id, name, data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		logger.Println("Querying record by name:", name)
		QueryByNameOfSimple(name)
	},
}

var idCmd = &cobra.Command{
	Use:   "id",
	Short: "Query data by id",
	Run: func(cmd *cobra.Command, args []string) {
		id := 0
		err := error(nil)
		if len(args) > 0 {
			id, err = strconv.Atoi(args[0])
			if err != nil {
				logger.Fatalln("parse id failed:", err)
			}
		} else {
			if gID > 0 {
				id = gID
			}
		}
		logger.Println("Querying record:", id)
		QueryByID(strconv.Itoa(id))
	},
}

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Query data by name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		logger.Println("Querying record:", name)
		QueryByName(name)
	},
}

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Query data by data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := args[0]
		logger.Println("Querying record:", data)
		QueryByData(data)
	},
}

var updateCmd = &cobra.Command{
	Use:   "u",
	Short: "Update data by id",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var updateNameCmd = &cobra.Command{
	Use:   "name",
	Short: "Update name of record by id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		UpdateName(strconv.Itoa(gID), name)
	},
}

var updateDataCmd = &cobra.Command{
	Use:   "data",
	Short: "Update data of record by id",
	Run: func(cmd *cobra.Command, args []string) {
		data := ""
		if len(args) > 0 {
			data = args[0]
		} else {
			// read data from stdin pipeline
			reader := bufio.NewReader(os.Stdin)
			inputBytes, err := io.ReadAll(reader)
			if err != nil {
				logger.Fatalln("read pipeline data failed:", err)
			} else {
				data = string(inputBytes)
			}
		}
		UpdateData(strconv.Itoa(gID), data)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "d",
	Short: "Delete data by id",
	Run: func(cmd *cobra.Command, args []string) {
		DelInfo(strconv.Itoa(gID))
	},
}

var countCmd = &cobra.Command{
	Use:   "c",
	Short: "Count data in database",
	Run: func(cmd *cobra.Command, args []string) {
		cnt := QueryCount()
		logger.Println("All record Count:", cnt)

		lastUpdate := QueryLastUpdate()
		lastUpdateTime := time.Unix(lastUpdate, 0).Format("2006-01-02 15:04:05")
		logger.Println("Last record update time:", lastUpdateTime)
	},
}
