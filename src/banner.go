/*
 * This file is part of the Hopia.
 * Copyright (c) 2023 Nathanne Isip
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

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
	fmt.Println("    Launches a Distributed Denial of")
	fmt.Println("                                 Service attack with a configuration.")

	greenPrn.Print("    hopia crack ")
	cyanPrn.Print("<hash> <config>")
	fmt.Println("  Brute force an encrypted text")
	fmt.Println("                                 using a text list.")

	greenPrn.Print("    hopia af ")
	cyanPrn.Print("<url> <list>")
	fmt.Println("        Find an admin login panel using")
	fmt.Println("                                 with a possible admin panel list.")

	greenPrn.Print("    hopia ps ")
	cyanPrn.Print("<url>")
	fmt.Println("               Scans an open port from a URL.")
}
