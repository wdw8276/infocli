package main

import (
	"fmt"
	"infocli/utils"
	"time"

	"github.com/bndr/gotabulate"
)

func decryptData(data string) string {
	if gKey == "" {
		return data
	}
	result, err := utils.Decrypt(data, gKey)
	if err != nil {
		return data
	}
	return result
}

func ShowTable(infos *[]Info, isSimple bool) {
	logger.Println("show detail:", gShowDetail)

	if isSimple {
		// simple format: ID / Name only
		var tb [][]interface{}
		for _, v := range *infos {
			row := []interface{}{v.ID, v.Name}
			tb = utils.Append(tb, row).([][]interface{})
		}
		t := gotabulate.Create(tb)
		t.SetHeaders([]string{"ID", "Name"})
		t.SetEmptyString("None")
		t.SetAlign("left")
		fmt.Println(t.Render("grid"))
	} else {
		var tb [][]interface{}
		for _, v := range *infos {
			var row []interface{}
			if gShowDetail {
				row = []interface{}{v.ID, v.Name, decryptData(v.Data), time.Unix(v.Created, 0).Format("2006-01-02 15:04:05"), time.Unix(v.Updated, 0).Format("2006-01-02 15:04:05")}
			} else {
				row = []interface{}{v.Name, decryptData(v.Data)}
			}
			tb = utils.Append(tb, row).([][]interface{})
		}
		t := gotabulate.Create(tb)
		if gShowDetail {
			t.SetHeaders([]string{"ID", "Name", "Data", "Created", "Updated"})
		} else {
			t.SetHeaders([]string{"Name", "Data"})
		}
		t.SetEmptyString("None")
		t.SetAlign("left")
		fmt.Println(t.Render("grid"))
	}
}
