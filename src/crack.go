package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func compareHash(hash string, dictString string, sumFunc func(data []byte) []byte) bool {
	sum := md5.Sum([]byte(dictString))
	return hash == fmt.Sprintf("%x", sum)
}

func check(hash string, algorithm string, dictString string) bool {
	switch algorithm {
	case "md5":
		return compareHash(hash, dictString, md5.New().Sum)

	case "sha1":
		return compareHash(hash, dictString, sha1.New().Sum)

	case "sha256":
		return compareHash(hash, dictString, sha256.New().Sum)
	}

	return false
}

func crack(hash string, configFile string) {
	greenPrn := color.New(color.FgGreen)
	redPrn := color.New(color.FgRed)
	cyanPrn := color.New(color.FgCyan)

	greenPrn.Print("  ‣ ")
	fmt.Print("Initiating ")
	greenPrn.Print("hash cracker")
	fmt.Println("...")

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

	dictFilename := ""
	hashAlg := ""

	keys := properties.Keys()
	for i := 0; i < len(keys); i++ {
		currentKey := keys[i]

		if currentKey == "algorithm" {
			algo, err := properties.Get(currentKey)

			if !err {
				greenPrn.Print("  ‣ ")
				redPrn.Print("Error ")
				fmt.Println("reading the hash algorithm from properties.")

				return
			}

			if algo != "md5" &&
				algo != "sha1" &&
				algo != "sha256" {
				greenPrn.Print("  ‣ ")
				redPrn.Print("Invalid ")
				fmt.Println("input algorithm on property file.")

				return
			}

			hashAlg += algo
		} else if currentKey == "dictionary" {
			dict, err := properties.Get(currentKey)

			if !err {
				greenPrn.Print("  ‣ ")
				redPrn.Print("Error ")
				fmt.Println("reading the dictionary value from properties.")

				return
			}

			dictFilename += dict
		} else {
			greenPrn.Print("  ‣ ")
			fmt.Print("Unknown property key: ")
			redPrn.Printf("%s", currentKey)
			fmt.Println(".")
		}
	}

	_, dictConfigExists := os.Stat(dictFilename)
	if os.IsNotExist((dictConfigExists)) {
		greenPrn.Print("  ‣ ")
		fmt.Print("Dictionary file doesn't ")
		redPrn.Print("exists")
		fmt.Printf(": %s\n", dictFilename)

		return
	}

	dictString, dictErr := readLinesFromFile(dictFilename)
	if dictErr != nil {
		greenPrn.Print("  ‣ ")
		fmt.Print("Error ")
		redPrn.Print("reading ")
		fmt.Printf("dictionary file: %s\n", dictFilename)

		return
	}

	for i := 0; i < len(dictString); i++ {
		currentDict := dictString[i]

		greenPrn.Printf("    [%d] ", i)
		fmt.Printf("Comparing '%s' to '%s' (", hash, currentDict)
		cyanPrn.Printf("%s", hashAlg)
		fmt.Println(").")

		if check(hash, hashAlg, currentDict) {
			greenPrn.Print("  ‣ ")
			fmt.Print("Matched: ")
			cyanPrn.Printf("%s\n", currentDict)

			return
		}
	}

	greenPrn.Print("  ‣ ")
	redPrn.Print("Failed ")
	fmt.Println("to find any match from dictionary.")
}
