package newbiecontest

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
	"strings"

	"github.com/gocolly/colly/v2"
)

const baseURL = "https://newbiecontest.org/"

func getProfileCrawling(id int) ProfileNewbiecontest {
	profile := ProfileNewbiecontest{}
	var err error = nil
	var rankTotal string

	profile.id = id
	rankTotal, profile.score, profile.level, profile.username, profile.avatar, err = getNewbiecontestCrawling(id)
	if err != nil {
		log.Fatal("Error", err)
	}
	profile.rank = strings.Split(rankTotal, "/")[0]

	nbTotalUsers, _ := strconv.Atoi(strings.Split(rankTotal, "/")[1])
	rank, _ := strconv.Atoi(profile.rank)
	profile.nbTotalUsers = fmt.Sprint(nbTotalUsers)
	profile.rankRelative = fmt.Sprintf("%03.1f%%", float64(rank)/float64(nbTotalUsers)*100.)

	fmt.Println(profile)
	return profile
}

func getNewbiecontestCrawling(id int) (rank string, score string, level string, username string, avatar image.Image, err error) {
	if id < 0 {
		rank, score, level, username = "0", "0", "0", "0"
		err = errors.New("Id is not valid")
	} else {
		c := colly.NewCollector()

		reRank := regexp.MustCompile("Position : [0-9]*/[0-9]*")
		reScore := regexp.MustCompile("Points : [0-9]*")
		reLevel := regexp.MustCompile("Royaume : .*")

		c.OnHTML("p.nospace", func(e *colly.HTMLElement) {
			match := reRank.FindString(e.Text)
			if match != "" {
				rank = match[11:]
			}

			match = reScore.FindString(e.Text)
			if match != "" {
				score = match[9:]
			}

			match = reLevel.FindString(e.Text)
			if match != "" {
				level = match[10:]
			}
		})

		c.OnHTML("h2 a span.bold", func(e *colly.HTMLElement) {
			username = e.Text
		})

		c.OnHTML(".memberHeaderBox img", func(e *colly.HTMLElement) {
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
				err = errors.New("Id doesn't exist")
			}
		})

		completeURL := baseURL + "index.php?page=info_membre&id=" + fmt.Sprint(id)
		c.Visit(completeURL)
	}
	return
}
