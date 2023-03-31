package main

import (
	"clliv-go/pkg/common"
	"fmt"
)

func init() {
	fmt.Println("before main init clliv-go")
}

func main() {
	common.TestChan()
}
