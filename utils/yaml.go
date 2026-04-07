package utils

import (
	"gopkg.in/yaml.v3"
)

func YamlLoads(yamlStr *string, mapData *map[string]interface{}) error {
	err := yaml.Unmarshal([]byte(*yamlStr), mapData)
	return err
}

func YamlDumps(mapData *map[string]interface{}) (string, error) {
	d, err := yaml.Marshal(mapData)
	return string(d), err
}

// // load
// data := make(map[string]interface{})
// if err := utils.YamlLoads(&str, &data); err != nil {
// 	utils.Log().Fatal("yaml config file error!")
// }
// utils.Log().Debug(data)

// // get map->list item
// for i, item := range data["list"].([]interface{}) {
// 	if s, fine := item.(string); fine {
// 		utils.Log().Debug("index=", i, ", value=", s)
// 	}
// }

// // dumps
// if dumps, err := utils.YamlDumps(&data); err == nil {
// 	utils.Log().Debug("dumps: ", dumps)
// }
