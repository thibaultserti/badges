package rootme

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg" // to deal with avatar
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
)

const baseURL = "https://root-me.org/"

func getProfileCrawling(username string) ProfileRootme {
	profile := ProfileRootme{}
	var err error = nil

	profile.username = username

	profile.rank, profile.score, profile.level, profile.nbChall, profile.avatar, err = getRootmeCrawling(username)
	if err != nil {
		log.Fatal("Error", err)
	}

	points, _ := strconv.Atoi(profile.score)
	profile.level = fmt.Sprint(computeLevel(points))

	// To avoid Too Many Requests Error
	time.Sleep(time.Second * 8)

	nbTotalUsers, _ := strconv.Atoi(getNbTotalUsersRootmeCrawling())
	rank, _ := strconv.Atoi(profile.rank)
	profile.nbTotalUsers = fmt.Sprint(nbTotalUsers)
	profile.rankRelative = fmt.Sprintf("%03.1f%%", float64(rank)/float64(nbTotalUsers)*100.)

	return profile
}

func getRootmeCrawling(username string) (rank string, score string, level string, nbChall string, avatar image.Image, err error) {
	if username == "" {
		rank, score, level, nbChall = "0", "0", "0", "0"
		err = errors.New("username is not valid")
	} else {
		c := colly.NewCollector()

		reNumber := regexp.MustCompile("[0-9]+")
		reRank := regexp.MustCompile("[0-9]+\nPlace")
		reScore := regexp.MustCompile("[0-9]+\nPoints")
		reNbchall := regexp.MustCompile("[0-9]+\nChallenges")

		c.OnHTML("div.medium-6 div.row div.small-6", func(e *colly.HTMLElement) {
			match := reRank.FindString(e.Text)
			if match != "" {
				rank = reNumber.FindString(match)
			}
			match = reScore.FindString(e.Text)
			if match != "" {
				score = reNumber.FindString(match)
			}

			match = reNbchall.FindString(e.Text)
			if match != "" {
				nbChall = reNumber.FindString(match)
			}

		})

		c.OnHTML("h1 img.logo_auteur", func(e *colly.HTMLElement) {
			urlAvatar := baseURL + e.Attr("src")
			response, _ := http.Get(urlAvatar)
			if err != nil {
				log.Fatal("Error", err)
			}
			defer response.Body.Close()

			imgBytes, _ := ioutil.ReadAll(response.Body)
			toto := bytes.NewReader(imgBytes)
			avatar, _, err = image.Decode(toto)

			if err != nil {
				log.Fatal("Error", err)
			}
		})

		c.OnResponse(func(r *colly.Response) {
			if r.StatusCode == 404 {
				rank = "0"
				err = errors.New("username doesn't exist")
			}
		})

		completeURL := baseURL + username
		c.Visit(completeURL)
	}
	return

}
func getNbTotalUsersRootmeCrawling() (nbTotalUsers string) {
	c := colly.NewCollector()
	first := true
	c.OnHTML("#counter div h1.counter_value", func(e *colly.HTMLElement) {
		if first {
			nbTotalUsers = e.Attr("data-count")
			first = false
			fmt.Println(nbTotalUsers)
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.Visit("https://pro.root-me.org")

	return
}
