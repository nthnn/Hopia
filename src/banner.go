package main

import (
	"fmt"

	"github.com/fatih/color"
)

func dumpHopiaBanner() {
	borderPrn := color.New(color.FgBlue)
	borderPrn.Println("                ┌─────────────────────────────────────┐")
	borderPrn.Print("                │")

	greenPrn := color.New(color.FgGreen)
	greenPrn.Print("      Hopia Wannabe-Hacker Tool")
	borderPrn.Println("      │")

	infoPrn := color.New(color.FgYellow)
	borderPrn.Print("                │")
	infoPrn.Print("           v0.0.1 - Beta             ")
	borderPrn.Println("│")
	borderPrn.Println("                └─────────────────────────────────────┘")

	fmt.Println()
}

func dumpUsage() {
	greenPrn := color.New(color.FgGreen)
	cyanPrn := color.New(color.FgCyan)

	fmt.Println("  Usage:")

	greenPrn.Print("    hopia ddos ")
	cyanPrn.Print("<url> <config>")
	fmt.Println("    Launches a Distributed Denial of\n                                 Service attack with a configuration.")

	greenPrn.Print("    hopia crack ")
	cyanPrn.Print("<hash> <config>")
	fmt.Println("  Brute force an encrypted text\n                                 using a text list.")

	greenPrn.Print("    hopia af ")
	cyanPrn.Print("<url> <config>")
	fmt.Println("      Find an admin login panel using\n                                 with a configuration and list.")

	greenPrn.Print("    hopia ps ")
	cyanPrn.Print("<url>")
	fmt.Println("               Scans an open port from a URL.")
}
