package main

import (
	"os"
)

func main() {
	dumpHopiaBanner()

	if len(os.Args) < 3 {
		dumpUsage()
		return
	}

	switch os.Args[1] {
	case "ddos":
		if len(os.Args) != 4 {
			dumpUsage()
			return
		}

		ddos(os.Args[2], os.Args[3])

	case "crack":
		if len(os.Args) != 4 {
			dumpUsage()
			return
		}

		crack(os.Args[2], os.Args[3])

	case "af":
		if len(os.Args) != 4 {
			dumpUsage()
			return
		}

		adminFinder(os.Args[2], os.Args[3])

	case "ps":
		break
	}
}
