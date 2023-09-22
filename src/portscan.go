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
