package utils

import (

)

func MapExistKey(m map[string]interface{}, k string) bool {
  if _, ok := m[k]; ok {
    return true
  }
  return false
}

func MapDeleteKey(m map[string]interface{}, k string) error {
  delete(m, k)  // no crash
  return nil
}
