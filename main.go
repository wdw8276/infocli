package main

// Time    :   2024-06-04 02:02:00 PM
// Author  :   diwen

import (
	"fmt"
	"infocli/utils"
	"os"
	"os/user"

	"gorm.io/gorm"
)

const ()

var (
	logger = utils.Logger // 日志对象

	// global flags 命令行参数
	gDbFile      string
	gID          int
	gShowDetail  bool
	gShowVersion bool
	gDebug       bool

	// global variables 数据库相关
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

	// default db file
	defaultDb := user.HomeDir + "/." + user.Username + ".db" // like: ~/.fish.db
	// logger.Println("default db file:", defaultDb)

	// add global flags
	rootCmd.PersistentFlags().StringVarP(&gDbFile, "file", "f", defaultDb, "database file")
	rootCmd.PersistentFlags().IntVarP(&gID, "id", "i", 0, "input record id")
	rootCmd.PersistentFlags().BoolVarP(&gShowDetail, "detail", "d", false, "show detail data")
	rootCmd.PersistentFlags().BoolVarP(&gShowVersion, "version", "v", false, "show version")
	rootCmd.PersistentFlags().BoolVarP(&gDebug, "debug", "D", false, "enable debug mode")

	// add sub commands
	// 打印版本信息
	rootCmd.AddCommand(versionCmd)

	// 初始化db
	rootCmd.AddCommand(initDbCmd)

	// 增加一条记录
	rootCmd.AddCommand(addCmd)

	// 查询数据 支持根据name、data、id 字段查询
	nameCmd.Flags().BoolVarP(&gShowDetail, "detail", "d", false, "show detail data")
	dataCmd.Flags().BoolVarP(&gShowDetail, "detail", "d", false, "show detail data")
	idCmd.Flags().BoolVarP(&gShowDetail, "detail", "d", false, "show detail data")
	queryCmd.AddCommand(nameCmd)
	queryCmd.AddCommand(dataCmd)
	queryCmd.AddCommand(idCmd)
	rootCmd.AddCommand(queryCmd)

	// 更新数据 包含根据id 和name 字段更新
	updateNameCmd.Flags().IntVarP(&gID, "id", "i", 0, "input record id")
	updateNameCmd.MarkFlagRequired("id")
	updateDataCmd.Flags().IntVarP(&gID, "id", "i", 0, "input record id")
	updateDataCmd.MarkFlagRequired("id")
	updateCmd.AddCommand(updateNameCmd)
	updateCmd.AddCommand(updateDataCmd)
	rootCmd.AddCommand(updateCmd)

	// 删除一条记录
	deleteCmd.Flags().IntVarP(&gID, "id", "i", 0, "input record id")
	deleteCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteCmd)

	// 统计行数
	rootCmd.AddCommand(countCmd)

	// 执行命令行
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	PrintGlobalFlags()
	os.Exit(0)
}
