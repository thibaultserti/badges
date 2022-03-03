package rootme

const levelmax = "legend"

var pointsLevel = [8]int{100, 500, 1976, 3454, 7999, 12687, 18226, 19500}
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
