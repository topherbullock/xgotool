package main

import (
	"fmt"
	"os"

	"github.com/topherbullock/xgotool/x11"
)

var (
	debug = true
)

func main() {
	display, err := x11.NewDisplay(os.Getenv("DISPLAY"))
	if err != nil {
		fmt.Errorf("err: %s", err.Error())
		os.Exit(1)
	}

	key := os.Args[1]
	display.Press(key)
}
