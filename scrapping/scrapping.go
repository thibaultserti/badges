package scrapping

import (
	"fmt"
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
	profile.username = username
	profile.rank = getRankCryptohackCrawling(username)
	profile.score = getScoreCryptohackCrawling(username)
	points, _ := strconv.Atoi(profile.score)
	profile.level = fmt.Sprint(computeLevel(points))

	nbTotalUsers, _ := strconv.Atoi(getNbTotalUsersCryptohackCrawling())
	fmt.Println(nbTotalUsers)
	rank, _ := strconv.Atoi(profile.rank)

	profile.rankRelative = fmt.Sprintf("%03.1f%%", float64(rank)/float64(nbTotalUsers)*100.)

	fmt.Printf("Profile: %#v\n", profile)
}

func getRankCryptohackCrawling(username string) (rank string) {
	c := colly.NewCollector()

	reRank := regexp.MustCompile("#[0-9]*")

	c.OnHTML(".userPoints p", func(e *colly.HTMLElement) {
		match := reRank.FindString(e.Text)
		if match != "" {
			rank = match[1:]
		}
	})
	completeURL := baseURL + "user/" + username
	c.Visit(completeURL)
	return
}

func getScoreCryptohackCrawling(username string) (score string) {
	c := colly.NewCollector()

	c.OnHTML("#userScore", func(e *colly.HTMLElement) {
		score = e.Text
	})

	completeURL := baseURL + "user/" + username
	c.Visit(completeURL)
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
