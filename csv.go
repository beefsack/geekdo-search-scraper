package main

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
