package newbiecontest

import (
	"log"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"github.com/thibaultserti/badges/common"
	"golang.org/x/image/font/gofont/goregular"
)

const width, height = 400, 160

// CreateNewbiecontestBadge creates the Newbiecontest badge
func CreateNewbiecontestBadge(id int, theme string, filename string) error {
	colorBG := new(common.Color)
	colorFG := new(common.Color)

	// crawl profile
	profile := getProfileCrawling(id)

	if theme == "dark" {
		colorBG.R, colorBG.G, colorBG.B = 0, 0, 0
		colorFG.R, colorFG.G, colorFG.B = 1, 1, 1
	} else if theme == "light" {
		colorBG.R, colorBG.G, colorBG.B = 1, 1, 1
		colorFG.R, colorFG.G, colorFG.B = 0, 0, 0
	}

	// load images
	logoNewbiecontest, err := gg.LoadImage("newbiecontest/images/newbiecontest-" + theme + ".png")
	if err != nil {
		log.Fatal(err)
	}

	userNewbiecontest := profile.avatar

	// load icons
	star, err := gg.LoadImage("icons/star.png")
	if err != nil {
		log.Fatal(err)
	}
	thunder, err := gg.LoadImage("icons/thunder.png")
	if err != nil {
		log.Fatal(err)
	}
	stats, err := gg.LoadImage("icons/stats.png")
	if err != nil {
		log.Fatal(err)
	}
	trophy, err := gg.LoadImage("icons/trophy.png")
	if err != nil {
		log.Fatal(err)
	}

	// resize images
	logoNewbiecontest = resize.Resize(0, 0.4*height, logoNewbiecontest, resize.Lanczos3)
	userNewbiecontest = resize.Resize(0, 0.6*height, userNewbiecontest, resize.Lanczos3)
	star = resize.Resize(width/20, 0, star, resize.Lanczos3)
	thunder = resize.Resize(width/20, 0, thunder, resize.Lanczos3)
	stats = resize.Resize(width/20, 0, stats, resize.Lanczos3)
	trophy = resize.Resize(width/20, 0, trophy, resize.Lanczos3)

	// setup fonts
	font, err := truetype.Parse(goregular.TTF)
	fontUsername := truetype.NewFace(font, &truetype.Options{Size: width / 16})
	fontOther := truetype.NewFace(font, &truetype.Options{Size: width / 28})

	dc := gg.NewContext(width, height)

	// set background color
	dc.SetRGB(colorBG.R, colorBG.G, colorBG.B)
	dc.Clear()

	// set font color
	dc.SetRGB(colorFG.R, colorFG.G, colorFG.B)
	// set username text
	dc.SetFontFace(fontUsername)
	dc.DrawStringAnchored(profile.username, width/2, height/4, 0.5, 0.5)

	// write other text
	dc.SetFontFace(fontOther)
	dc.DrawStringAnchored("Level: "+profile.level, width/5, 0.85*height, 0.4, 0.5)
	dc.DrawImageAnchored(thunder, 0.1*width, 0.85*height, 0.5, 0.5)
	dc.DrawStringAnchored(profile.score+" points", width/2, 0.45*height, 0.4, 0.5)
	dc.DrawImageAnchored(star, 0.4*width, 0.45*height, 0.5, 0.5)
	dc.DrawStringAnchored(profile.rank+"/"+profile.nbTotalUsers, width/2, 0.6*height, 0.4, 0.5)
	dc.DrawImageAnchored(trophy, 0.4*width, 0.6*height, 0.5, 0.5)
	dc.DrawStringAnchored("TOP "+profile.rankRelative, width/2, 0.75*height, 0.4, 0.5)
	dc.DrawImageAnchored(stats, 0.4*width, 0.75*height, 0.5, 0.5)

	// draw images
	dc.DrawImageAnchored(logoNewbiecontest, 6*width/7, height/3, 0.5, 0.5)
	dc.DrawImageAnchored(userNewbiecontest, width/5, height/2, 0.5, 0.6)

	err = dc.SavePNG(filename) // save it
	return err
}
