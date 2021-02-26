package rootme

import (
	"image"
)

// ProfileRootme describe a user profile on Rootme
type ProfileRootme struct {
	username     string
	score        string
	nbChall      string
	level        string
	rank         string
	rankRelative string
	nbTotalUsers string
	avatar       image.Image
}
