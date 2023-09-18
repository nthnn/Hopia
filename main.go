package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func dumpUsage() {
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

	cyanPrn := color.New(color.FgCyan)
	fmt.Println("\n  Usage:")

	greenPrn.Print("    hopia ddos ")
	cyanPrn.Print("<config>")
	fmt.Println("   Launches a Distributed Denial of\n                          Service attack with a configuration.")

	greenPrn.Print("    hopia crack ")
	cyanPrn.Print("<config>")
	fmt.Println("  Brute force an encrypted text\n                          using a text list.")

	greenPrn.Print("    hopia af ")
	cyanPrn.Print("<config>")
	fmt.Println("     Find an admin login panel using\n                          with a configuration and list.")

	greenPrn.Print("    hopia ps ")
	cyanPrn.Print("<url>")
	fmt.Println("        Scans an open port from a URL.")
}

func main() {
	if len(os.Args) == 1 {
		dumpUsage()
		return
	}
}
