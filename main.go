package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println(color.RedString("Usage: <mode> <target>"));
		return
	}

	

}
