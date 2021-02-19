package scrapping

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/gocolly/colly/v2"
)

type Profile struct {
	username     string
	score        string
	level        string
	rank         string
	rankRelative string
}

const baseURL = "https://cryptohack.org/"

func GetProfileCrawling(username string) {
	profile := Profile{}
	var err error = nil

	profile.username = username

	profile.rank, err = getRankCryptohackCrawling(username)
	if err != nil {
		log.Fatal("Error", err)
	}
	profile.score, err = getScoreCryptohackCrawling(username)
	if err != nil {
		log.Fatal("Error", err)
	}
	points, _ := strconv.Atoi(profile.score)
	profile.level = fmt.Sprint(computeLevel(points))

	nbTotalUsers, _ := strconv.Atoi(getNbTotalUsersCryptohackCrawling())
	fmt.Println(nbTotalUsers)
	rank, _ := strconv.Atoi(profile.rank)

	profile.rankRelative = fmt.Sprintf("%03.1f%%", float64(rank)/float64(nbTotalUsers)*100.)

	fmt.Printf("Profile: %#v\n", profile)
}

func getRankCryptohackCrawling(username string) (rank string, err error) {
	if username == "" {
		rank = "0"
		err = errors.New("Username is empty")
	} else {
		c := colly.NewCollector()

		reRank := regexp.MustCompile("#[0-9]*")

		c.OnResponse(func(r *colly.Response) {
			if r.StatusCode == 404 {
				rank = "0"
				err = errors.New("Username doesn't exist")
			}
		})

		c.OnHTML(".userPoints p", func(e *colly.HTMLElement) {
			match := reRank.FindString(e.Text)
			if match != "" {
				rank = match[1:]
			}
		})
		completeURL := baseURL + "user/" + username
		c.Visit(completeURL)
	}
	return
}

func getScoreCryptohackCrawling(username string) (score string, err error) {
	if username == "" {
		score = "0"
		err = errors.New("Username is empty")
	} else {
		c := colly.NewCollector()

		c.OnHTML("#userScore", func(e *colly.HTMLElement) {
			score = e.Text
		})

		c.OnResponse(func(r *colly.Response) {
			if r.StatusCode == 404 {
				score = "0"
				err = errors.New("Username doesn't exist")
			}
		})
		completeURL := baseURL + "user/" + username
		c.Visit(completeURL)
	}
	return
}

func getNbTotalUsersCryptohackCrawling() (nbTotalUsers string) {
	c := colly.NewCollector()

	reNbTotalUsers := regexp.MustCompile("Users: [0-9]*")

	c.OnHTML(".scoreboardStats", func(e *colly.HTMLElement) {
		match := reNbTotalUsers.FindString(e.Text)
		if match != "" {
			nbTotalUsers = match[7:]
		}
	})
	completeURL := baseURL + "scoreboard/"
	c.Visit(completeURL)
	return
}
