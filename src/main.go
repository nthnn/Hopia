package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func dumpUsage() {
	greenPrn := color.New(color.FgGreen)
	cyanPrn := color.New(color.FgCyan)

	fmt.Println("  Usage:")

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
	dumpHopiaBanner()

	if len(os.Args) != 3 {
		dumpUsage()
		return
	}

	switch os.Args[1] {
	case "ddos":
		ddos(os.Args[2])
		break

	default:
		fmt.Println()
	}
}
