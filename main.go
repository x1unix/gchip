package main
/**
Package main contains core emulator logic.

Based on: http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/
 */

import (
	"gchip/emulator"
	"fmt"
)

func main() {
	m := emulator.CreateMachine()

	fmt.Println(m.Test())
}


