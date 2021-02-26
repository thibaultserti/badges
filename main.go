package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/thibaultserti/badges/cryptohack"
	"github.com/thibaultserti/badges/newbiecontest"
	"github.com/thibaultserti/badges/rootme"
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

	availableWebsites := map[string]([]string){"cryptohack": {"dark", "light"}, "newbiecontest": {"dark", "light"}, "rootme": {"dark", "light"}}

	website := flag.String("website", "cryptohack", "Specify the challenge website from which you want to create the image (Available: cryptohack)")
	username := flag.String("username", "", "Specify the username or the id (depending on the website)")
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
	if *website == "newbiecontest" {
		id, err := strconv.Atoi(*username)
		if err != nil {
			log.Fatal(err)
		}
		newbiecontest.CreateNewbiecontestBadge(id, *theme)
	} else if *website == "cryptohack" {
		cryptohack.CreateCryptohackBadge(*username, *theme)
	} else if *website == "rootme" {
		rootme.CreateRootmeBadge(*username, *theme)
	}
}
