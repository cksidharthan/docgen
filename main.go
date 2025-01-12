//go:generate go run generate-asset.go

package main

import (
	"fmt"
	"github.com/cksidharthan/docgen/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println("error running docgen: ", err)
	}
}
