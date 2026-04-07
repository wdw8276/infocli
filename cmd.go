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
	// check show version
	if gShowVersion {
		logger.Println(gName, gVersion)
		// os.Exit(0)
	}
}

// init cobra -- 在cobra初始化时调用
func init() {
	cobra.OnInitialize(initConfig)
}

// cobra root command
var rootCmd = &cobra.Command{
	Use:   "infocli",
	Short: "A simple tool to store and query data, based on sqlite3.\nAuthor: diwen @Copyright 2024",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
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

// 初始化db
var initDbCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database",
	Run: func(cmd *cobra.Command, args []string) {
		// 初始化db
		InitDB()
	},
}

// 增加一条记录
var addCmd = &cobra.Command{
	Use:   "a",
	Short: "Add a record, first as name, second as data, or read from pipeline",
	Args:  cobra.MinimumNArgs(1), // 要求输入至少一个参数
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		data := ""
		if len(args) > 1 {
			data = args[1]
		} else {
			// 从管道标准输入读取数据
			reader := bufio.NewReader(os.Stdin)
			// 一次性读取所有输入内容直到遇到文件末尾（EOF）
			inputBytes, err := io.ReadAll(reader)
			if err != nil {
				logger.Fatalln("read pipeline data failed:", err)
			} else {
				data = string(inputBytes)
			}
		}
		logger.Println("Adding record:", name, data)

		// 调用orm写入数据库
		SaveInfo(name, data)
	},
}

// 查询数据，简单搜索 name 返回id和name
var queryCmd = &cobra.Command{
	Use:   "q",
	Short: "Query data by input name, or using subcommands to query by id, name, data",
	Args:  cobra.ExactArgs(1), // 要求输入一个参数
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		logger.Println("Querying record by name:", name)

		// 调用orm查询数据库
		QueryByNameOfSimple(name)
	},
}

// 查询数据，根据id
var idCmd = &cobra.Command{
	Use:   "id",
	Short: "Query data by id",
	// Args:  cobra.ExactArgs(1), // 要求输入一个参数
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
		// 调用orm查询数据库
		QueryByID(strconv.Itoa(id))
	},
}

// 查询数据，根据name
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Query data by name",
	Args:  cobra.ExactArgs(1), // 要求输入一个参数
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		logger.Println("Querying record:", name)
		// 调用orm查询数据库
		QueryByName(name)
	},
}

// 查询数据，根据data
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Query data by data",
	Args:  cobra.ExactArgs(1), // 要求输入一个参数
	Run: func(cmd *cobra.Command, args []string) {
		data := args[0]
		logger.Println("Querying record:", data)
		// 调用orm查询数据库
		QueryByData(data)
	},
}

// 根据id更新数据
var updateCmd = &cobra.Command{
	Use:   "u",
	Short: "Update data by id",
	Run: func(cmd *cobra.Command, args []string) {
		// logger.Println("Updating record... -h for help.")
		// logger.Println("ID:", gID)
	},
}

// 更新数据name
var updateNameCmd = &cobra.Command{
	Use:   "name",
	Short: "Update name of record by id",
	Args:  cobra.ExactArgs(1), // 要求输入一个参数
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		// logger.Println("Updating record, id:", gID, "name:", name)
		// 调用orm更新数据库
		UpdateName(strconv.Itoa(gID), name)
	},
}

// 更新数据data
var updateDataCmd = &cobra.Command{
	Use:   "data",
	Short: "Update data of record by id",
	Run: func(cmd *cobra.Command, args []string) {
		data := ""
		if len(args) > 0 {
			data = args[0]
		} else {
			// 从管道标准输入读取数据
			reader := bufio.NewReader(os.Stdin)
			// 一次性读取所有输入内容直到遇到文件末尾（EOF）
			inputBytes, err := io.ReadAll(reader)
			if err != nil {
				logger.Fatalln("read pipeline data failed:", err)
			} else {
				data = string(inputBytes)
			}
		}
		// logger.Println("Updating record, id:", gID, "data:", data)
		// 调用orm更新数据库
		UpdateData(strconv.Itoa(gID), data)
	},
}

// 删除数据
var deleteCmd = &cobra.Command{
	Use:   "d",
	Short: "Delete data by id",
	Run: func(cmd *cobra.Command, args []string) {
		// 调用orm删除数据库
		DelInfo(strconv.Itoa(gID))
	},
}

// 统计数据库中数据条数
var countCmd = &cobra.Command{
	Use:   "c",
	Short: "Count data in database",
	Run: func(cmd *cobra.Command, args []string) {
		// 调用orm统计数据库
		cnt := QueryCount()
		logger.Println("All record Count:", cnt)

		// 查询最后一条记录的更新时间戳
		lastUpdate := QueryLastUpdate()
		// 转换为可读时间戳
		lastUpdateTime := time.Unix(lastUpdate, 0).Format("2006-01-02 15:04:05")
		logger.Println("Last record update time:", lastUpdateTime)
	},
}
