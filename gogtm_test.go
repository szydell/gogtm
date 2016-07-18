package gogtm_test

import (
  "gogtm"
  "fmt"
)

func ExampleSet() {
  err := gogtm.Set("^test","aaa")
  if err != nil {
    fmt.Println("gogtm.Set failed: ", err.Error())
  }
}
