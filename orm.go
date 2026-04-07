package main

import (
	"fmt"
	"infocli/utils"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// table orm 表结构
type Info struct {
	ID      uint64 `gorm:"primaryKey"`
	Created int64  `gorm:"autoCreateTime"` // Use unix seconds as creating time
	Updated int64  `gorm:"autoUpdateTime"` // Use unix milli seconds as updating time
	Name    string `gorm:"unique"`
	Data    string
}

// 初始化数据库表结构
func InitDB() {
	if !utils.FileExist(gDbFile) {
		logger.Println("create db file:", gDbFile)
	} else {
		logger.Println("load db file:", gDbFile)
	}
	gDb, gErr = gorm.Open(sqlite.Open(gDbFile), &gorm.Config{})
	if gErr != nil {
		logger.Fatalln(gErr)
	}

	// 自动创建或修改数据库表结构
	gErr = gDb.AutoMigrate(&Info{})
	if gErr != nil {
		logger.Fatalln(gErr)
	} else {
		logger.Println("Auto refresh table ok")
	}
}

// 保存一条记录
func SaveInfo(name string, data string) error {
	// connect to db
	InitDB()

	var info = Info{
		Name: name,
		Data: data,
	}

	// insert to db table
	gRet = gDb.Create(&info)
	if gRet.Error != nil {
		logger.Fatalln(gRet.Error)
	}
	logger.Println("create record:", info.Name, "ok")
	return nil
}

// 删除一条记录
func DelInfo(id string) error {
	logger.Println("delete a record..")

	logger.Println("ID:", id)
	if !utils.IsNumeric(id) {
		logger.Fatalln("not a numeric, abort")
	}

	// connect to db
	InitDB()

	// delete record of id
	gRet = gDb.Delete(&Info{}, id)
	if gRet.Error != nil {
		logger.Fatalln(gRet.Error)
	}
	if gRet.RowsAffected < 1 {
		logger.Println("no this record of id:", id)
	} else {
		logger.Println("delete id:", id, "record ok")
	}
	return nil
}

// 更新一条记录的name字段
func UpdateName(id string, name string) error {
	logger.Println("update a record..")

	logger.Println("ID:", id)
	if !utils.IsNumeric(id) {
		logger.Fatalln("not a numeric, abort")
	}
	logger.Println("Name:", name)

	// connect to db
	InitDB()

	// update record of id
	gRet = gDb.Model(&Info{}).Where("id = ?", id).Update("name", name)
	if gRet.Error != nil {
		logger.Fatalln(gRet.Error)
	}
	if gRet.RowsAffected < 1 {
		logger.Println("no this record of id:", id)
	} else {
		logger.Println("update id:", id, "record's name ok")
	}
	return nil
}

// 更新一条记录的data字段
func UpdateData(id string, data string) error {
	logger.Println("update a record..")

	logger.Println("ID:", id)
	if !utils.IsNumeric(id) {
		logger.Fatalln("not a numeric, abort")
	}
	logger.Println("Data:", data)

	// connect to db
	InitDB()

	// update record of id
	gRet = gDb.Model(&Info{}).Where("id = ?", id).Update("data", data)
	if gRet.Error != nil {
		logger.Fatalln(gRet.Error)
	}
	if gRet.RowsAffected < 1 {
		logger.Println("no this record of id:", id)
	} else {
		logger.Println("update id:", id, "record's data ok")
	}
	return nil
}

// 通过id字段搜索
func QueryByID(id string) error {
	// connect to db
	InitDB()
	// query record by id
	var infos []Info
	gRet = gDb.Where("id =?", id).Find(&infos)
	if gRet.RowsAffected < 1 {
		logger.Println("not find any records")
		return nil
	} else {
		// show records
		ShowTable(&infos, false)
	}
	return nil
}

// 通过name字段搜索 - 简单模式 只显示id和name字段
func QueryByNameOfSimple(name string) error {
	// connect to db
	InitDB()

	// query record by name
	var infos []Info
	like := fmt.Sprintf("%%%s%%", name)

	gRet = gDb.Where("name LIKE ?", like).Find(&infos)
	if gRet.RowsAffected < 1 {
		logger.Println("not find any records")
		return nil
	}

	// show records
	ShowTable(&infos, true)
	return nil
}

// 通过name字段搜索
func QueryByName(name string) error {
	// connect to db
	InitDB()

	// query record by name
	var infos []Info
	like := fmt.Sprintf("%%%s%%", name)

	gRet = gDb.Where("name LIKE ?", like).Find(&infos)
	if gRet.RowsAffected < 1 {
		logger.Println("not find any records")
		return nil
	}

	// show records
	ShowTable(&infos, false)
	return nil
}

// 通过data字段搜索
func QueryByData(data string) error {
	// connect to db
	InitDB()

	// build orm sql
	var infos []Info
	like := fmt.Sprintf("%%%s%%", data)

	gRet = gDb.Where("data LIKE?", like).Find(&infos)
	if gRet.RowsAffected < 1 {
		logger.Println("not find any records")
		return nil
	}

	// show records
	ShowTable(&infos, false)
	return nil
}

// 查询表中的记录总数
func QueryCount() int64 {
	// connect to db
	InitDB()

	// 查询表中的记录总数
	var count int64
	_ = gDb.Model(&Info{}).Count(&count)
	return count
}

// 查询最后一条记录的更新时间戳
func QueryLastUpdate() int64 {
	// connect to db
	InitDB()

	// 查询最后一条记录的更新时间戳
	var lastUpdate int64
	_ = gDb.Model(&Info{}).Order("updated DESC").Limit(1).Pluck("updated", &lastUpdate).Error
	return lastUpdate
}
