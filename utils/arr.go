package utils

import (
  "reflect"
  // "fmt"
)

func AppendInt(slice []int, item int) []int {
  return append(slice, item)
}

func AppendString(slice []string, item string) []string {
  return append(slice, item)
}

func AppendByte(slice []byte, item byte) []byte {
  return append(slice, item)
}

// NOTE: return interface{}, to switch to array, using r.([]int) or r.([]string)
// l := []string{"1", "2"}
// l = utils.Append(l, "3").([]string)  // same array pointer address

func Append(slice interface{}, item interface{}) interface{} {
  v := reflect.ValueOf(slice)
  v = reflect.Append(v, reflect.ValueOf(item))
  return v.Interface()
}

// NOTE: 字符串和int数组，删除和替换，返回新的数组
func DelStringArrayItem(slice []string, item string) []string {
  l := make([]string, 0)
  for _, v := range slice {
    if v != item {
      l = append(l, v)
    }
  }
  return l
}

func ReplaceStringArrayItem(slice []string, find string, item string) []string {
  l := make([]string, 0)
  for _, v := range slice {
    if v == find {
      l = append(l, item)
    } else {
      l = append(l, v)
    }
  }
  return l
}

func DelIntArrayItem(slice []int, item int) []int {
  l := make([]int, 0)
  for _, v := range slice {
    if v != item {
      l = append(l, v)
    }
  }
  return l
}

func ReplaceIntArrayItem(slice []int, find int, item int) []int {
  l := make([]int, 0)
  for _, v := range slice {
    if v == find {
      l = append(l, item)
    } else {
      l = append(l, v)
    }
  }
  return l
}
