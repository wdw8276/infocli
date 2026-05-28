package main

import (
	"fmt"
	"infocli/utils"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Info is the ORM model for the info table
type Info struct {
	ID      uint64 `gorm:"primaryKey"`
	Created int64  `gorm:"autoCreateTime"` // unix seconds
	Updated int64  `gorm:"autoUpdateTime"` // unix milliseconds
	Name    string `gorm:"unique"`
	Data    string
}

func InitDB() {
	if gDb != nil {
		return
	}
	if !utils.FileExist(gDbFile) {
		logger.Println("create db file:", gDbFile)
	} else {
		logger.Println("load db file:", gDbFile)
	}
	gDb, gErr = gorm.Open(sqlite.Open(gDbFile), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if gErr != nil {
		logger.Fatalln(gErr)
	}

	gErr = gDb.AutoMigrate(&Info{})
	if gErr != nil {
		logger.Fatalln(gErr)
	} else {
		logger.Println("Auto refresh table ok")
	}
}

func SaveInfo(name string, data string) error {
	if gKey != "" {
		encrypted, err := utils.Encrypt(data, gKey)
		if err != nil {
			logger.Fatalln("encrypt failed:", err)
		}
		data = encrypted
	}

	var info = Info{
		Name: name,
		Data: data,
	}

	gRet = gDb.Create(&info)
	if gRet.Error != nil {
		if strings.Contains(gRet.Error.Error(), "UNIQUE constraint failed") {
			logger.Println("record already exists:", name)
			return nil
		}
		logger.Fatalln(gRet.Error)
	}
	logger.Println("create record:", info.Name, "ok")
	return nil
}

func DelInfo(id string) error {
	logger.Println("delete a record..")
	logger.Println("ID:", id)
	if !utils.IsNumeric(id) {
		logger.Fatalln("not a numeric, abort")
	}

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

func UpdateName(id string, name string) error {
	logger.Println("update a record..")
	logger.Println("ID:", id)
	if !utils.IsNumeric(id) {
		logger.Fatalln("not a numeric, abort")
	}
	logger.Println("Name:", name)

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

func UpdateData(id string, data string) error {
	logger.Println("update a record..")
	logger.Println("ID:", id)
	if !utils.IsNumeric(id) {
		logger.Fatalln("not a numeric, abort")
	}
	logger.Println("Data:", data)

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

func QueryByID(id string) error {
	var infos []Info
	gRet = gDb.Where("id =?", id).Find(&infos)
	if gRet.RowsAffected < 1 {
		logger.Println("not find any records")
		return nil
	}
	ShowTable(&infos, false)
	return nil
}

// QueryByNameOfSimple queries by name and shows only ID and Name columns
func QueryByNameOfSimple(name string) error {
	var infos []Info
	like := fmt.Sprintf("%%%s%%", name)

	gRet = gDb.Where("name LIKE ?", like).Find(&infos)
	if gRet.RowsAffected < 1 {
		logger.Println("not find any records")
		return nil
	}
	ShowTable(&infos, true)
	return nil
}

func QueryByName(name string) error {
	var infos []Info
	like := fmt.Sprintf("%%%s%%", name)

	gRet = gDb.Where("name LIKE ?", like).Find(&infos)
	if gRet.RowsAffected < 1 {
		logger.Println("not find any records")
		return nil
	}
	ShowTable(&infos, false)
	return nil
}

func QueryByData(data string) error {
	var infos []Info
	like := fmt.Sprintf("%%%s%%", data)

	gRet = gDb.Where("data LIKE?", like).Find(&infos)
	if gRet.RowsAffected < 1 {
		logger.Println("not find any records")
		return nil
	}
	ShowTable(&infos, false)
	return nil
}

func QueryCount() int64 {
	var count int64
	_ = gDb.Model(&Info{}).Count(&count)
	return count
}

func QueryLastUpdate() int64 {
	var lastUpdate int64
	_ = gDb.Model(&Info{}).Order("updated DESC").Limit(1).Pluck("updated", &lastUpdate).Error
	return lastUpdate
}
