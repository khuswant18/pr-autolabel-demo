package main

import (
	"fmt"
	"os"

	"github.com/khuswant18/pr-autolabel-demo/core/scan"
)

func main() {
	target := "cluster"
	if len(os.Args) > 1 {
		target = os.Args[1]
	}
	fmt.Println(scan.Run(target))
}
