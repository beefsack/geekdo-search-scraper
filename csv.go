package main

import "fmt"

const (
	columnID = iota
	columnType
	columnName
	columnYear
	columnURL
	columnThumbnail
	columnImage
	columnMinPlayers
	columnMaxPlayers
	columnPlayingTime
	columnUsersRated
	columnAverage
	columnBayesAverage
	columnStdDev
	columnMedian
	columnOwned
	columnTrading
	columnWanting
	columnWishing
	columnNumComments
	columnNumWeights
	columnAverageWeight
	columnRank
	columnRatings1
	columnRatings2
	columnRatings3
	columnRatings4
	columnRatings5
	columnRatings6
	columnRatings7
	columnRatings8
	columnRatings9
	columnRatings10
	// Player age poll
	columnAgePoll2
	columnAgePoll3
	columnAgePoll4
	columnAgePoll5
	columnAgePoll6
	columnAgePoll8
	columnAgePoll10
	columnAgePoll12
	columnAgePoll14
	columnAgePoll16
	columnAgePoll18
	columnAgePoll21
	// Language dependence poll
	columnLanguagePoll1
	columnLanguagePoll2
	columnLanguagePoll3
	columnLanguagePoll4
	columnLanguagePoll5
	// Suggested player poll
	columnPlayerPoll1Best
	columnPlayerPoll1Recommended
	columnPlayerPoll1NotRecommended
	columnPlayerPoll2Best
	columnPlayerPoll2Recommended
	columnPlayerPoll2NotRecommended
	columnPlayerPoll3Best
	columnPlayerPoll3Recommended
	columnPlayerPoll3NotRecommended
	columnPlayerPoll4Best
	columnPlayerPoll4Recommended
	columnPlayerPoll4NotRecommended
	columnPlayerPoll5Best
	columnPlayerPoll5Recommended
	columnPlayerPoll5NotRecommended
	columnPlayerPoll6Best
	columnPlayerPoll6Recommended
	columnPlayerPoll6NotRecommended
	columnPlayerPoll7Best
	columnPlayerPoll7Recommended
	columnPlayerPoll7NotRecommended
	columnPlayerPoll8Best
	columnPlayerPoll8Recommended
	columnPlayerPoll8NotRecommended
	columnPlayerPoll9Best
	columnPlayerPoll9Recommended
	columnPlayerPoll9NotRecommended
	columnPlayerPoll10Best
	columnPlayerPoll10Recommended
	columnPlayerPoll10NotRecommended
)

const (
	columnPlayerPollFormat               = "Player Poll %s %s"
	columnPlayerPollBestFormat           = "Player Poll %s Best"
	columnPlayerPollRecommendedFormat    = "Player Poll %s Recommended"
	columnPlayerPollNotRecommendedFormat = "Player Poll %s Not Recommended"
	columnAgePollFormat                  = "Age Poll %s"
	columnLanguagePollFormat             = "Language Poll %s"
)

var columns = []int{
	columnID,
	columnType,
	columnName,
	columnYear,
	columnURL,
	columnThumbnail,
	columnImage,
	columnMinPlayers,
	columnMaxPlayers,
	columnPlayingTime,
	columnUsersRated,
	columnAverage,
	columnBayesAverage,
	columnStdDev,
	columnMedian,
	columnOwned,
	columnTrading,
	columnWanting,
	columnWishing,
	columnNumComments,
	columnNumWeights,
	columnAverageWeight,
	columnRank,
	columnRatings1,
	columnRatings2,
	columnRatings3,
	columnRatings4,
	columnRatings5,
	columnRatings6,
	columnRatings7,
	columnRatings8,
	columnRatings9,
	columnRatings10,
	// Player age poll
	columnAgePoll2,
	columnAgePoll3,
	columnAgePoll4,
	columnAgePoll5,
	columnAgePoll6,
	columnAgePoll8,
	columnAgePoll10,
	columnAgePoll12,
	columnAgePoll14,
	columnAgePoll16,
	columnAgePoll18,
	columnAgePoll21,
	// Language dependence poll
	columnLanguagePoll1,
	columnLanguagePoll2,
	columnLanguagePoll3,
	columnLanguagePoll4,
	columnLanguagePoll5,
	// Suggested player poll
	columnPlayerPoll1Best,
	columnPlayerPoll1Recommended,
	columnPlayerPoll1NotRecommended,
	columnPlayerPoll2Best,
	columnPlayerPoll2Recommended,
	columnPlayerPoll2NotRecommended,
	columnPlayerPoll3Best,
	columnPlayerPoll3Recommended,
	columnPlayerPoll3NotRecommended,
	columnPlayerPoll4Best,
	columnPlayerPoll4Recommended,
	columnPlayerPoll4NotRecommended,
	columnPlayerPoll5Best,
	columnPlayerPoll5Recommended,
	columnPlayerPoll5NotRecommended,
	columnPlayerPoll6Best,
	columnPlayerPoll6Recommended,
	columnPlayerPoll6NotRecommended,
	columnPlayerPoll7Best,
	columnPlayerPoll7Recommended,
	columnPlayerPoll7NotRecommended,
	columnPlayerPoll8Best,
	columnPlayerPoll8Recommended,
	columnPlayerPoll8NotRecommended,
	columnPlayerPoll9Best,
	columnPlayerPoll9Recommended,
	columnPlayerPoll9NotRecommended,
	columnPlayerPoll10Best,
	columnPlayerPoll10Recommended,
	columnPlayerPoll10NotRecommended,
}

var columnTitles = map[int]string{
	columnID:            "ID",
	columnType:          "Type",
	columnName:          "Name",
	columnYear:          "Year",
	columnURL:           "URL",
	columnThumbnail:     "Thumbnail",
	columnImage:         "Image",
	columnMinPlayers:    "Min Players",
	columnMaxPlayers:    "Max Players",
	columnPlayingTime:   "Playing Time",
	columnUsersRated:    "Users Rated",
	columnAverage:       "Average",
	columnBayesAverage:  "Bayes Average",
	columnStdDev:        "Std Dev",
	columnMedian:        "Median",
	columnOwned:         "Owned",
	columnTrading:       "Trading",
	columnWanting:       "Wanting",
	columnWishing:       "Wishing",
	columnNumComments:   "Num Comments",
	columnNumWeights:    "Num Weights",
	columnAverageWeight: "Average Weight",
	columnRank:          "Rank",
	columnRatings1:      "Ratings 1",
	columnRatings2:      "Ratings 2",
	columnRatings3:      "Ratings 3",
	columnRatings4:      "Ratings 4",
	columnRatings5:      "Ratings 5",
	columnRatings6:      "Ratings 6",
	columnRatings7:      "Ratings 7",
	columnRatings8:      "Ratings 8",
	columnRatings9:      "Ratings 9",
	columnRatings10:     "Ratings 10",
	// Player age poll
	columnAgePoll2:  fmt.Sprintf(columnAgePollFormat, "2"),
	columnAgePoll3:  fmt.Sprintf(columnAgePollFormat, "3"),
	columnAgePoll4:  fmt.Sprintf(columnAgePollFormat, "4"),
	columnAgePoll5:  fmt.Sprintf(columnAgePollFormat, "5"),
	columnAgePoll6:  fmt.Sprintf(columnAgePollFormat, "6"),
	columnAgePoll8:  fmt.Sprintf(columnAgePollFormat, "8"),
	columnAgePoll10: fmt.Sprintf(columnAgePollFormat, "10"),
	columnAgePoll12: fmt.Sprintf(columnAgePollFormat, "12"),
	columnAgePoll14: fmt.Sprintf(columnAgePollFormat, "14"),
	columnAgePoll16: fmt.Sprintf(columnAgePollFormat, "16"),
	columnAgePoll18: fmt.Sprintf(columnAgePollFormat, "18"),
	columnAgePoll21: fmt.Sprintf(columnAgePollFormat, "21 and up"),
	// Language dependence poll
	columnLanguagePoll1: fmt.Sprintf(columnLanguagePollFormat, "No necessary in-game text"),
	columnLanguagePoll2: fmt.Sprintf(columnLanguagePollFormat, "Some necessary text - easily memorized or small crib sheet"),
	columnLanguagePoll3: fmt.Sprintf(columnLanguagePollFormat, "Moderate in-game text - needs crib sheet or paste ups"),
	columnLanguagePoll4: fmt.Sprintf(columnLanguagePollFormat, "Extensive use of text - massive conversion needed to be playable"),
	columnLanguagePoll5: fmt.Sprintf(columnLanguagePollFormat, "Unplayable in another language"),
	// Suggested player poll
	columnPlayerPoll1Best:            fmt.Sprintf(columnPlayerPollBestFormat, "1"),
	columnPlayerPoll1Recommended:     fmt.Sprintf(columnPlayerPollRecommendedFormat, "1"),
	columnPlayerPoll1NotRecommended:  fmt.Sprintf(columnPlayerPollNotRecommendedFormat, "1"),
	columnPlayerPoll2Best:            fmt.Sprintf(columnPlayerPollBestFormat, "2"),
	columnPlayerPoll2Recommended:     fmt.Sprintf(columnPlayerPollRecommendedFormat, "2"),
	columnPlayerPoll2NotRecommended:  fmt.Sprintf(columnPlayerPollNotRecommendedFormat, "2"),
	columnPlayerPoll3Best:            fmt.Sprintf(columnPlayerPollBestFormat, "3"),
	columnPlayerPoll3Recommended:     fmt.Sprintf(columnPlayerPollRecommendedFormat, "3"),
	columnPlayerPoll3NotRecommended:  fmt.Sprintf(columnPlayerPollNotRecommendedFormat, "3"),
	columnPlayerPoll4Best:            fmt.Sprintf(columnPlayerPollBestFormat, "4"),
	columnPlayerPoll4Recommended:     fmt.Sprintf(columnPlayerPollRecommendedFormat, "4"),
	columnPlayerPoll4NotRecommended:  fmt.Sprintf(columnPlayerPollNotRecommendedFormat, "4"),
	columnPlayerPoll5Best:            fmt.Sprintf(columnPlayerPollBestFormat, "5"),
	columnPlayerPoll5Recommended:     fmt.Sprintf(columnPlayerPollRecommendedFormat, "5"),
	columnPlayerPoll5NotRecommended:  fmt.Sprintf(columnPlayerPollNotRecommendedFormat, "5"),
	columnPlayerPoll6Best:            fmt.Sprintf(columnPlayerPollBestFormat, "6"),
	columnPlayerPoll6Recommended:     fmt.Sprintf(columnPlayerPollRecommendedFormat, "6"),
	columnPlayerPoll6NotRecommended:  fmt.Sprintf(columnPlayerPollNotRecommendedFormat, "6"),
	columnPlayerPoll7Best:            fmt.Sprintf(columnPlayerPollBestFormat, "7"),
	columnPlayerPoll7Recommended:     fmt.Sprintf(columnPlayerPollRecommendedFormat, "7"),
	columnPlayerPoll7NotRecommended:  fmt.Sprintf(columnPlayerPollNotRecommendedFormat, "7"),
	columnPlayerPoll8Best:            fmt.Sprintf(columnPlayerPollBestFormat, "8"),
	columnPlayerPoll8Recommended:     fmt.Sprintf(columnPlayerPollRecommendedFormat, "8"),
	columnPlayerPoll8NotRecommended:  fmt.Sprintf(columnPlayerPollNotRecommendedFormat, "8"),
	columnPlayerPoll9Best:            fmt.Sprintf(columnPlayerPollBestFormat, "9"),
	columnPlayerPoll9Recommended:     fmt.Sprintf(columnPlayerPollRecommendedFormat, "9"),
	columnPlayerPoll9NotRecommended:  fmt.Sprintf(columnPlayerPollNotRecommendedFormat, "9"),
	columnPlayerPoll10Best:           fmt.Sprintf(columnPlayerPollBestFormat, "10"),
	columnPlayerPoll10Recommended:    fmt.Sprintf(columnPlayerPollRecommendedFormat, "10"),
	columnPlayerPoll10NotRecommended: fmt.Sprintf(columnPlayerPollNotRecommendedFormat, "10"),
}

// Header gets the CSV header.
func Header() []string {
	l := len(columns)
	header := make([]string, l)
	for i, c := range columns {
		header[i] = columnTitles[c]
	}
	return header
}

// Row converts values into a CSV row.
func Row(values map[int]string) []string {
	l := len(columns)
	row := make([]string, l)
	for i, c := range columns {
		row[i] = values[c]
	}
	return row
}

// HeaderIndex takes a header string and returns the index, if it exists.
func HeaderIndex(header string) (int, bool) {
	for i, h := range columnTitles {
		if header == h {
			return i, true
		}
	}
	return 0, false
}
