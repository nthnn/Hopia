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
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/fatih/color"
)

func adminFinder(urlStr string, listFile string) {
	greenPrn := color.New(color.FgGreen)
	redPrn := color.New(color.FgRed)
	cyanPrn := color.New(color.FgCyan)

	greenPrn.Print("  ‣ ")
	fmt.Print("Initiating ")
	greenPrn.Print("admin finder")
	fmt.Println(" tool...")

	host, errHost := url.ParseRequestURI(urlStr)
	if errHost != nil {
		greenPrn.Print("  ‣ ")
		redPrn.Print("Error ")
		fmt.Printf("parsing host: %s\n", errHost.Error())

		return
	}

	hostname := host.Scheme + "://" + host.Host
	_, listExists := os.Stat(listFile)

	if os.IsNotExist(listExists) {
		greenPrn.Print("  ‣ ")
		fmt.Print("List file ")
		redPrn.Print("not found")
		fmt.Println(".")

		return
	}

	paths, err := readLinesFromFile(listFile)
	if err != nil {
		greenPrn.Print("  ‣ ")
		redPrn.Print("Error ")
		fmt.Printf("reading list file: %s.\n", err.Error())

		return
	}

	var wg sync.WaitGroup
	foundPaths := []string{}

	for i := 0; i < len(paths); i++ {
		currentPath := paths[i]
		wg.Add(1)

		go func(iterNum int) {
			defer wg.Done()

			client := &http.Client{}
			target := hostname + "/" + currentPath

			req := createRequest(target)
			resp, errClient := client.Do(req)

			if errClient != nil {
				greenPrn.Print("  ‣ ")
				redPrn.Print("Error")
				fmt.Printf(": %s.\n", err.Error())

				return
			} else if resp.StatusCode == 200 {
				foundPaths = append(foundPaths, target)
			}

			greenPrn.Printf("    [%d] ", iterNum)
			fmt.Printf("Requested: %s (", target)
			cyanPrn.Printf("%d", resp.StatusCode)
			fmt.Println(").")
		}(i)
	}

	wg.Wait()
	if len(foundPaths) == 0 {
		greenPrn.Print("  ‣ ")
		redPrn.Print("No ")
		fmt.Print("login panel ")
		redPrn.Println("was found.")

		return
	}

	greenPrn.Print("  ‣ Found ")
	fmt.Println("the following admin panel(s):")

	for i := 0; i < len(foundPaths); i++ {
		fmt.Print("    - ")
		cyanPrn.Println(foundPaths[i])
	}
}
