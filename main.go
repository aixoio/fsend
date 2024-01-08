package main

import (
	"fmt"
	"os"

	"github.com/aixoio/fsend/decoder"
	"github.com/aixoio/fsend/encoder"
	"github.com/aixoio/fsend/helper"
	"github.com/fatih/color"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(color.RedString("Usage: <mode>"));
		return
	}
	
	if args[1] == "-h" || args[1] == "--help" {
		helper.Help();
		return
	}

	if len(args) < 3 {
		fmt.Println(color.RedString("Usage: <mode> <target>"));
		return
	}

	if args[1] == "-t" {
		encoder.Start(args[2]);
		return
	}

	if args[1] == "-r" {
		decoder.Start(args[2]);
		return
	}

	fmt.Println(color.RedString("Unknow input mode: %s", args[1]))

}
