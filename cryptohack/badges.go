package cryptohack

import (
	"log"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/font/gofont/goregular"
)

const width, height = 400, 200

// CreateCryptohackBadge creates the cryptohack badge
func CreateCryptohackBadge(username string, theme string) error {
	colorBG := new(Color)
	colorFG := new(Color)

	if theme == "dark" {
		colorBG.r, colorBG.g, colorBG.b = 12./255., 18./255., 33./255.
		colorFG.r, colorFG.g, colorFG.b = 1, 1, 1
	} else if theme == "light" {
		colorBG.r, colorBG.g, colorBG.b = 1, 1, 1
		colorFG.r, colorFG.g, colorFG.b = 240./255., 78./255., 35./255.
	}

	// load images
	logoCryptohack, err := gg.LoadImage("cryptohack/images/cryptohack.png")
	if err != nil {
		log.Fatal(err)
	}

	userCryptohack, err := gg.LoadImage("cryptohack/images/user.png")
	if err != nil {
		log.Fatal(err)
	}

	// resize images
	logoCryptohack = resize.Resize(0, 0.4*height, logoCryptohack, resize.Lanczos3)
	userCryptohack = resize.Resize(0, 0.8*height, userCryptohack, resize.Lanczos3)

	// crawl profile
	profile := getProfileCrawling(username)

	// setup fonts
	font, err := truetype.Parse(goregular.TTF)
	fontUsername := truetype.NewFace(font, &truetype.Options{Size: width / 16})
	fontOther := truetype.NewFace(font, &truetype.Options{Size: width / 32})

	dc := gg.NewContext(width, height)

	// set background color
	dc.SetRGB(colorBG.r, colorBG.g, colorBG.b)
	dc.Clear()

	// set font color
	dc.SetRGB(colorFG.r, colorFG.g, colorFG.b)
	// set username text
	dc.SetFontFace(fontUsername)
	dc.DrawStringAnchored(profile.username, width/2, height/4, 0.5, 0.5)

	// write other text
	dc.SetFontFace(fontOther)
	dc.DrawStringAnchored("Level: "+profile.level, width/5, 0.85*height, 0.5, 0.5)
	dc.DrawStringAnchored(profile.score+" points", width/2, 0.4*height, 0.5, 0.5)
	dc.DrawStringAnchored(profile.rank+"/"+profile.nbTotalUsers, width/2, 0.5*height, 0.5, 0.5)
	dc.DrawStringAnchored("TOP "+profile.rankRelative, width/2, 0.6*height, 0.5, 0.5)

	// draw images
	dc.DrawImageAnchored(logoCryptohack, 6*width/7, height/3, 0.5, 0.5)
	dc.DrawImageAnchored(userCryptohack, width/5, height/2, 0.5, 0.6)

	err = dc.SavePNG("cryptohack.png") // save it
	return err
}
