package newbiecontest

import (
	"image"
)

// ProfileNewbiecontest describe a user profile on Newbiecontest
type ProfileNewbiecontest struct {
	id           int
	username     string
	score        string
	level        string
	rank         string
	rankRelative string
	nbTotalUsers string
	avatar       image.Image
}
