package rootme

const levelmax = "legend"

var pointsLevel = [8]int{100, 500, 2040, 3581, 8322, 13212, 18990, 20340}
var nameLevel = [8]string{"visitor", "curious", "trainee", "insider", "enthusiast", "hacker", "elite", "legend"}

func computeLevel(score int) (level string) {
	level = levelmax
	for lvl, points := range pointsLevel {
		if points > score {
			level = nameLevel[lvl]
			break
		}
	}
	return
}
