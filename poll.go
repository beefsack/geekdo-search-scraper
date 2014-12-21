package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/beefsack/go-geekdo"
)

func parseSingleResultSetPoll(poll geekdo.ThingPoll, headerFormat string) (
	map[int]string,
	error,
) {
	if poll.ResultSets == nil || len(poll.ResultSets) == 0 ||
		poll.ResultSets[0].Results == nil ||
		len(poll.ResultSets[0].Results) == 0 {
		return nil, errors.New("there are no poll results to parse")
	}
	results := map[int]string{}
	for _, r := range poll.ResultSets[0].Results {
		if h, ok := HeaderIndex(fmt.Sprintf(headerFormat, r.Value)); ok {
			results[h] = strconv.Itoa(r.NumVotes)
		}
	}
	return results, nil
}

func parseNumPlayersPoll(poll geekdo.ThingPoll) map[int]string {
	results := map[int]string{}

	if poll.ResultSets == nil {
		return results
	}
	for _, rs := range poll.ResultSets {
		if rs.Results == nil {
			continue
		}
		for _, r := range rs.Results {
			if h, ok := HeaderIndex(fmt.Sprintf(
				columnPlayerPollFormat, rs.NumPlayers, r.Value)); ok {
				results[h] = strconv.Itoa(r.NumVotes)
			}
		}
	}
	return results
}
