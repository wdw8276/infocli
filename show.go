package main

import (
	"fmt"
	"infocli/utils"
	"time"

	"github.com/bndr/gotabulate"
)

// 使用表格打印记录
func ShowTable(infos *[]Info, isSimple bool) {
	// 是否需要显示详细信息
	logger.Println("show detail:", gShowDetail)

	// 打印简单格式 id/name
	if isSimple {
		var tb [][]interface{}

		for _, v := range *infos {
			var row = []interface{}{v.ID, v.Name}
			tb = utils.Append(tb, row).([][]interface{})
		}

		// 设置表头 并打印
		t := gotabulate.Create(tb)
		t.SetHeaders([]string{"ID", "Name"})
		t.SetEmptyString("None")
		t.SetAlign("left")
		fmt.Println(t.Render("grid"))
	} else {
		// 打印普通格式 id/name/data
		var tb [][]interface{}

		for _, v := range *infos {
			var row []interface{}
			// 打印详细信息
			if gShowDetail {
				row = []interface{}{v.ID, v.Name, v.Data, time.Unix(v.Created, 0).Format("2006-01-02 15:04:05"), time.Unix(v.Updated, 0).Format("2006-01-02 15:04:05")}
			} else {
				row = []interface{}{v.Name, v.Data}
			}
			tb = utils.Append(tb, row).([][]interface{})
		}

		// 设置表头
		t := gotabulate.Create(tb)
		if gShowDetail {
			t.SetHeaders([]string{"ID", "Name", "Data", "Created", "Updated"})
		} else {
			t.SetHeaders([]string{"Name", "Data"})
		}

		// 打印
		t.SetEmptyString("None")
		t.SetAlign("left")
		fmt.Println(t.Render("grid"))
	}
}
