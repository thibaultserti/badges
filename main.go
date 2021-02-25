package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thibaultserti/badges/cryptohack"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	availableWebsites := map[string]([]string){"cryptohack": {"dark", "light"}}

	website := flag.String("website", "cryptohack", "Specify the challenge website from which you want to create the image (Available: cryptohack)")
	username := flag.String("username", "", "Specify the username")
	theme := flag.String("theme", "dark", "Specify the theme")

	flag.Parse()

	if len(availableWebsites[*website]) == 0 {
		fmt.Println("Website not supported")
		os.Exit(1)
	}

	if *username == "" {
		fmt.Println("Username not defined")
		os.Exit(1)
	}

	if !contains(availableWebsites[*website], *theme) {
		fmt.Println("Theme not supported for this website")
		os.Exit(1)
	}

	cryptohack.CreateCryptohackBadge(*username, *theme)
}
