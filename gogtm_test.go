package gogtm_test

import (
	"fmt"

	"github.com/szydell/gogtm"
)

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

func ExampleOrder() {
	data, err := gogtm.Order("^test", "")
	if err != nil {
		fmt.Println("gogtm.Order failed", err.Error())
	}
	rev, err := gogtm.Order("^test", "-1")
	if err != nil {
		fmt.Println("gogtm.Order with the reverse direction flag failed", err.Error())
	}
	fmt.Sprintf("gogtm.Order returned %s, in the reverse direction returned %s", data, rev)
}

func ExampleSet() {
	err := gogtm.Set("^test", "aaa")
	if err != nil {
		fmt.Println("gogtm.Set failed: ", err.Error())
	}
}

func ExampleZKill() {
	err := gogtm.ZKill("^test")
	if err != nil {
		fmt.Println("gogtm.ZKill failed: ", err.Error())
	}
}
