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
