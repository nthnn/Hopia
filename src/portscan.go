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
	"strconv"
	"sync"

	"github.com/fatih/color"
)

func portScan(urlHost string) {
	greenPrn := color.New(color.FgGreen)
	redPrn := color.New(color.FgRed)

	greenPrn.Print("  ‣ ")
	fmt.Print("Initiating ")
	greenPrn.Print("port scanner")
	fmt.Println("...")

	host, err := url.ParseRequestURI(urlHost)
	if err != nil {
		greenPrn.Print("  ‣ ")
		redPrn.Print("Invalid ")
		fmt.Printf("input URL: %s\n", urlHost)

		return
	}

	var wg sync.WaitGroup
	hostname := host.Scheme + "://" + host.Host
	foundPorts := []int{}

	for i := 1; i < 65536; i++ {
		wg.Add(1)

		go func(iterNum int) {
			client := &http.Client{}
			req := createRequest(hostname + ":" + strconv.Itoa(iterNum))
			resp, err := client.Do(req)

			defer wg.Done()
			if err == nil &&
				((resp.StatusCode >= 200 && resp.StatusCode <= 207) ||
					(resp.StatusCode >= 400 && resp.StatusCode <= 451)) {
				foundPorts = append(foundPorts, iterNum)
			}
		}(i)
	}

	wg.Wait()
	if len(foundPorts) == 0 {
		greenPrn.Print("  ‣ ")
		fmt.Print("No port(s) was ")
		redPrn.Print("found")
		fmt.Print(".")

		return
	}

	greenPrn.Print("  ‣ ")
	fmt.Print("Ports ")
	greenPrn.Print("found")
	fmt.Print(":\n    ")
	fmt.Println(foundPorts)

	greenPrn.Print("  ‣ ")
	fmt.Print("Scanning ")
	greenPrn.Print("done")
	fmt.Println("!")
}
