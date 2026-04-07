package utils

import (
	jsoniter "github.com/json-iterator/go"
)

// NOTE: example
// jsonStr := `{"num":6.13, "name":"fish"}`
// data := make(map[string]interface{})
//
// if e := utils.JsonLoads(&jsonStr, &data); e == nil {
// 	logger.Println("JsonLoads:", data)
// }
//
// if s, e := utils.JsonDumps(&data); e == nil {
// 	logger.Println("JsonDumps:", s)
// }

var (
	gJson = jsoniter.ConfigCompatibleWithStandardLibrary
)

func JsonLoads(jsonStr *string, mapData *map[string]interface{}) error {
	jsonByteArr := []byte(*jsonStr)
	return gJson.Unmarshal(jsonByteArr, mapData)
}

func JsonDumps(mapData *map[string]interface{}) (string, error) {
	s, err := gJson.Marshal(mapData)
	return string(s), err
}
