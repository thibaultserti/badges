package scrapping

import (
	"strconv"
	"testing"
)

const userTest = "thibaultserti"
const userNotExisting = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

func TestGetRankCryptohackCrawling(t *testing.T) {
	rank, err := getRankCryptohackCrawling(userTest)
	_, err = strconv.Atoi(rank)

	if err != nil {
		t.Error("Test failed; rank is not a int", rank)
	}
}

func TestGetScoreCryptohackCrawling(t *testing.T) {
	score, err := getScoreCryptohackCrawling(userTest)
	_, err = strconv.Atoi(score)
	if err != nil {
		t.Error("Test failed: score is not a int", score)
	}
}

func TestGetRankNullUserCryptohackCrawling(t *testing.T) {
	_, err := getRankCryptohackCrawling("")
	if err == nil {
		t.Error("Test failed: no error for null user")
	}
}

func TestGetScoreNullUserCryptohackCrawling(t *testing.T) {
	_, err := getScoreCryptohackCrawling("")
	if err == nil {
		t.Error("Test failed: no error for null user")
	}
}

func TestGetRankAbsentUserCryptohackCrawling(t *testing.T) {
	_, err := getRankCryptohackCrawling(userNotExisting)

	if err != nil {
		t.Error("Test failed: no error for absent user")
	}
}

func TestGetScoreAbsentUserCryptohackCrawling(t *testing.T) {
	_, err := getScoreCryptohackCrawling(userNotExisting)
	if err != nil {
		t.Error("Test failed: no error for absent user")
	}
}

func TestGetNbTotalUsersCryptohackCrawling(t *testing.T) {
	nbTotalUsers, err := strconv.Atoi(getNbTotalUsersCryptohackCrawling())
	if err != nil {
		t.Error("Testt failed: nbTotalUsers is not a int", nbTotalUsers)
	}

}
