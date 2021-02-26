package rootme

const levelmax = "elite"

var pointsLevel = [5]int{3500, 7000, 10500, 14000, 17500}
var nameLevel = [5]string{"newbie", "lamer", "programmer", "hacker", "elite"}

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
