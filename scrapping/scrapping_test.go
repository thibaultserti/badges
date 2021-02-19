package scrapping

import (
	"strconv"
	"testing"
)

const userTest = "thibaultserti"
const userNotExisting = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

func TestGetRankCryptohackCrawling(t *testing.T) {
	_, err := strconv.Atoi(getRankCryptohackCrawling(userTest))

	if err != nil {
		t.Error("Not a int")
	}
}

func TestGetScoreCryptohackCrawling(t *testing.T) {
	_, err := strconv.Atoi(getScoreCryptohackCrawling(userTest))
	if err != nil {
		t.Error("Not a int")
	}
}

func TestGetNbTotalUsersCryptohackCrawling(t *testing.T) {
	_, err := strconv.Atoi(getNbTotalUsersCryptohackCrawling())
	if err != nil {
		t.Error("Not a int")
	}

}
