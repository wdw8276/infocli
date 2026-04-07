package utils

import (
  "flag"
)

func TypeOfVar(v interface{}) string {
  switch v.(type) {
  case int:
    return "int"
  case bool:
    return "bool"
  case string:
    return "string"
  case float32:
    return "float32"
  case float64:
    return "float64"
  case *[]int:
    return "*[]int"
  case *[]string:
    return "*[]string"
  case *[]byte:
    return "*[]byte"
  default:
    return "unknow"
  }
}


func SetFlagString(flagShortName string, defaultValue string, tips string) *string {
  return flag.String(flagShortName, defaultValue, tips)
}

func SetFlagBool(flagShortName string, defaultValue bool, tips string) *bool {
  return flag.Bool(flagShortName, defaultValue, tips)
}

func SetFlagInt(flagShortName string, defaultValue int, tips string) *int {
  return flag.Int(flagShortName, defaultValue, tips)
}

func FlagShowHelp()  {
  flag.Parse()
}
