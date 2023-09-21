package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func createRequest(client http.Client, url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	return req
}

func ddos(configFile string) {
	greenPrn := color.New(color.FgGreen)
	redPrn := color.New(color.FgRed)

	greenPrn.Print("  ‣ ")
	fmt.Print("Initiating ")
	greenPrn.Print("DDoS")
	fmt.Println(" attack...")

	_, configExists := os.Stat(configFile)
	if os.IsNotExist((configExists)) {
		greenPrn.Print("  ‣ ")
		redPrn.Print("Error ")
		fmt.Printf("reading config file: %s.\n", configFile)

		return
	}
}
