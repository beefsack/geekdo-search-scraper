package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beefsack/go-geekdo"
)

const (
	advSearchRetries            = 5
	advSearchWaitBetweenRetries = time.Minute
	thingRetries                = 5
	thingWaitBetweenRetries     = time.Minute
	ratingsRetries              = 5
	ratingsWaitBetweenRetries   = time.Minute
)

// AdvSearch scrapes an advanced search page for more powerful searching and
// ordering.
func AdvSearch(
	client *geekdo.Client,
	url string,
	logger *log.Logger,
) ([]geekdo.SearchCollectionItem, error) {
	errors := 0
	for errors < advSearchRetries {
		coll, err := client.AdvSearch(url)
		if err == nil {
			return coll, err
		}
		errors++
		logger.Printf(
			"AdvSearch failed for URL '%s', waiting %s before retry %d, %v",
			url,
			advSearchWaitBetweenRetries,
			errors,
			err,
		)
		time.Sleep(advSearchWaitBetweenRetries)
	}
	return nil, fmt.Errorf("failed after %d retries", advSearchRetries)
}

// Thing queries a thing of a kind and an id.
func Thing(
	client *geekdo.Client,
	kind string,
	id int,
	logger *log.Logger,
) (geekdo.Things, error) {
	errors := 0
	for errors < thingRetries {
		things, err := client.Thing(kind, id, geekdo.ThingOptions{
			Stats: true,
		})
		if err == nil {
			return things, err
		}
		errors++
		logger.Printf(
			"Thing failed for %s %d, waiting %s before retry %d, %v",
			kind,
			id,
			thingWaitBetweenRetries,
			errors,
			err,
		)
		time.Sleep(thingWaitBetweenRetries)
	}
	return geekdo.Things{}, fmt.Errorf("failed after %d retries", thingRetries)
}

// Ratings gets the rating counts (ratings from 1 to 10) for a thing.
func Ratings(
	client *geekdo.Client,
	id int,
	logger *log.Logger,
) (map[int]int, error) {
	errors := 0
	for errors < ratingsRetries {
		ratings, err := client.Ratings(id)
		if err == nil {
			return ratings, err
		}
		errors++
		logger.Printf(
			"Ratings failed for %d, waiting %s before retry %d, %v",
			id,
			ratingsWaitBetweenRetries,
			errors,
			err,
		)
		time.Sleep(ratingsWaitBetweenRetries)
	}
	return nil, fmt.Errorf("failed after %d retries", ratingsRetries)
}
