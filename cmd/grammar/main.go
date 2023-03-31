package main

import (
	"fmt"

	"github.com/cl1111/clliv-go/pkg/common"
)

func init() {
	fmt.Println("before main init clliv-go")
}

func main() {
	common.TestChan()
}
