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

func ExampleGet() {
	data, err := gogtm.Get("^test", "ababababababababa")
	if err != nil {
		fmt.Println("gogtm.Get failed: ", err.Error())
	}
	if data == "ababababababababa" {
		fmt.Println("Global ^test empty, default value returned")
	} else {
		fmt.Println("Value of global ^test is: " + data)
	}
}

func ExampleKill() {
	err := gogtm.Kill("^test")
	if err != nil {
		fmt.Println("gogtm.Kill failed: ", err.Error())
	}
}
