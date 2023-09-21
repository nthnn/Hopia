package main

import (
	"fmt"

	"github.com/fatih/color"
)

func dumpHopiaBanner() {
	borderPrn := color.New(color.FgBlue)
	borderPrn.Println("        ┌─────────────────────────────────────┐")
	borderPrn.Print("        │")

	greenPrn := color.New(color.FgGreen)
	greenPrn.Print("      Hopia Wannabe-Hacker Tool")
	borderPrn.Println("      │")

	infoPrn := color.New(color.FgYellow)
	borderPrn.Print("        │")
	infoPrn.Print("           v0.0.1 - Beta             ")
	borderPrn.Println("│")
	borderPrn.Println("        └─────────────────────────────────────┘")

	fmt.Println()
}
