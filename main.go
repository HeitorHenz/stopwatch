package main

import (
	"fmt"
	"stopwatch/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
