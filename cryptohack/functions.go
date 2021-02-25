package cryptohack

const levelmax = 25

var pointsLevel = [levelmax]int{0, 10, 20, 40, 80, 160, 320, 640, 1280, 1920, 2560, 3200, 3840, 4480, 5120, 5760, 6400, 7040, 7680, 8320, 8960, 9600, 10240, 10880, 11520}

func computeLevel(score int) (level int) {
	level = levelmax
	for lvl, points := range pointsLevel {
		if points > score {
			level = lvl
			break
		}
	}
	return
}

// Color represents the color with rgb value between 0 and 1
type Color struct {
	r float64
	g float64
	b float64
}
