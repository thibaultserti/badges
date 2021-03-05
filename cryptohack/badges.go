package cryptohack

import (
	"log"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"github.com/thibaultserti/badges/common"
	"golang.org/x/image/font/gofont/goregular"
)

const width, height = 400, 160

// CreateCryptohackBadge creates the cryptohack badge
func CreateCryptohackBadge(username string, theme string, filename string) error {
	colorBG := new(common.Color)
	colorFG := new(common.Color)

	if theme == "dark" {
		colorBG.R, colorBG.G, colorBG.B = 12./255., 18./255., 33./255.
		colorFG.R, colorFG.G, colorFG.B = 1, 1, 1
	} else if theme == "light" {
		colorBG.R, colorBG.G, colorBG.B = 1, 1, 1
		colorFG.R, colorFG.G, colorFG.B = 240./255., 78./255., 35./255.
	}

	// load images
	logoCryptohack, err := gg.LoadImage("cryptohack/images/cryptohack.png")
	if err != nil {
		log.Fatal(err)
	}
	nameCryptohack, err := gg.LoadImage("cryptohack/images/cryptohack-name-" + theme + ".png")
	if err != nil {
		log.Fatal(err)
	}

	userCryptohack, err := gg.LoadImage("cryptohack/images/user.png")
	if err != nil {
		log.Fatal(err)
	}

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
	logoCryptohack = resize.Resize(0, 0.4*height, logoCryptohack, resize.Lanczos3)
	nameCryptohack = resize.Resize(0, 0.175*height, nameCryptohack, resize.Lanczos3)
	userCryptohack = resize.Resize(0, 0.8*height, userCryptohack, resize.Lanczos3)
	star = resize.Resize(width/20, 0, star, resize.Lanczos3)
	thunder = resize.Resize(width/20, 0, thunder, resize.Lanczos3)
	stats = resize.Resize(width/20, 0, stats, resize.Lanczos3)
	trophy = resize.Resize(width/20, 0, trophy, resize.Lanczos3)

	// crawl profile
	profile := getProfileCrawling(username)

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
	dc.DrawImageAnchored(logoCryptohack, 6*width/7, height/3, 0.5, 0.5)
	dc.DrawImageAnchored(nameCryptohack, 4*width/5, 0.9*height, 0.5, 0.5)
	dc.DrawImageAnchored(userCryptohack, width/5, height/2, 0.5, 0.6)

	err = dc.SavePNG(filename) // save it
	return err
}
