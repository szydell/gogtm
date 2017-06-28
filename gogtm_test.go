package gogtm_test

import (
	"fmt"

	"github.com/szydell/gogtm"
)

func ExampleSet() {
	err := gogtm.Set("^test", "aaa")
	if err != nil {
		fmt.Println("gogtm.Set failed: ", err.Error())
	}
}
