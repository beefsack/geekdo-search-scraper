package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/beefsack/go-geekdo"
)

func main() {
	stderr := log.New(os.Stderr, "", 0)
	urls := os.Args[1:]
	if len(urls) == 0 {
		stderr.Fatalf("You must specify one or more URLs")
	}
	client := geekdo.NewClient()
	w := csv.NewWriter(os.Stdout)
	if err := w.Write(Header()); err != nil {
		stderr.Fatalf("Error writing header, %s", err)
	}
	w.Flush()
	for _, u := range urls {
		coll, err := AdvSearch(client, u, stderr)
		if err != nil {
			stderr.Fatalf("Error fetching search page, %s", err)
		}
		for _, c := range coll {
			var (
				ratings geekdo.ThingStatsRatings
				rank    geekdo.ThingStatsRank
			)
			// API request for more data
			items, err := Thing(client, c.Type, c.ID, stderr)
			if err != nil {
				stderr.Fatalf("Error querying API for %d, %s", c.ID, err)
			}
			if len(items.Items) == 0 {
				stderr.Fatalf("Could not fetch anything for %d", c.ID)
			}
			thing := items.Items[0]
			if len(thing.Statistics.Ratings) > 0 {
				ratings = thing.Statistics.Ratings[0]
				if len(ratings.Ranks) > 0 {
					rank = ratings.Ranks[0]
				}
			}
			// Scrape ratings
			userRatings, err := Ratings(client, thing.ID, stderr)
			if err != nil {
				stderr.Fatalf("Could not fetch ratings for %d, %s", c.ID, err)
			}
			// Build row
			values := map[int]string{
				columnID:            fmt.Sprintf("%d", thing.ID),
				columnType:          c.Type,
				columnName:          thing.Names[0].Value,
				columnYear:          fmt.Sprintf("%d", thing.YearPublished.Value),
				columnURL:           c.URL,
				columnThumbnail:     thing.Thumbnail,
				columnImage:         thing.Image,
				columnMinPlayers:    fmt.Sprintf("%d", thing.MinPlayers.Value),
				columnMaxPlayers:    fmt.Sprintf("%d", thing.MaxPlayers.Value),
				columnPlayingTime:   fmt.Sprintf("%d", thing.PlayingTime.Value),
				columnUsersRated:    fmt.Sprintf("%d", ratings.UsersRated.Value),
				columnAverage:       fmt.Sprintf("%f", ratings.Average.Value),
				columnBayesAverage:  fmt.Sprintf("%f", ratings.BayesAverage.Value),
				columnStdDev:        fmt.Sprintf("%f", ratings.StdDev.Value),
				columnMedian:        fmt.Sprintf("%f", ratings.Median.Value),
				columnOwned:         fmt.Sprintf("%d", ratings.Owned.Value),
				columnTrading:       fmt.Sprintf("%d", ratings.Trading.Value),
				columnWanting:       fmt.Sprintf("%d", ratings.Wanting.Value),
				columnWishing:       fmt.Sprintf("%d", ratings.Wishing.Value),
				columnNumComments:   fmt.Sprintf("%d", ratings.NumComments.Value),
				columnNumWeights:    fmt.Sprintf("%d", ratings.NumWeights.Value),
				columnAverageWeight: fmt.Sprintf("%f", ratings.AverageWeight.Value),
				columnRank:          fmt.Sprintf("%d", rank.Value),
				columnRatings1:      fmt.Sprintf("%d", userRatings[1]),
				columnRatings2:      fmt.Sprintf("%d", userRatings[2]),
				columnRatings3:      fmt.Sprintf("%d", userRatings[3]),
				columnRatings4:      fmt.Sprintf("%d", userRatings[4]),
				columnRatings5:      fmt.Sprintf("%d", userRatings[5]),
				columnRatings6:      fmt.Sprintf("%d", userRatings[6]),
				columnRatings7:      fmt.Sprintf("%d", userRatings[7]),
				columnRatings8:      fmt.Sprintf("%d", userRatings[8]),
				columnRatings9:      fmt.Sprintf("%d", userRatings[9]),
				columnRatings10:     fmt.Sprintf("%d", userRatings[10]),
			}
			// Polls
			if thing.Polls != nil {
				for _, poll := range thing.Polls {
					results := map[int]string{}
					switch poll.Name {
					case "suggested_numplayers":
					case "suggested_playerage":
						results, err = parseSingleResultSetPoll(poll,
							columnAgePollFormat)
						if err != nil {
							stderr.Printf("Unable to parse age poll, %v", err)
							continue
						}
					case "language_dependence":
						results, err = parseSingleResultSetPoll(poll,
							columnLanguagePollFormat)
						if err != nil {
							stderr.Printf("Unable to parse language poll, %v", err)
							continue
						}
					}
					for c, v := range results {
						values[c] = v
					}
				}
			}
			if err := w.Write(Row(values)); err != nil {
				stderr.Fatalf("Error writing row, %s", err)
			}
		}
		w.Flush()
	}
}
