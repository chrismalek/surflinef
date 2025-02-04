package main

import (
	"fmt"
	"net/url"

	"github.com/chrismalek/surflinef/v2"
)

func main() {
	getConditions()
	getTides()
	getTaxonomy()
	getWave()
}

func getConditions() {
	bu, err := url.Parse(surflinef.ConditionsBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	c := surflinef.Client{BaseURL: bu}

	cq := surflinef.ConditionsQuery{
		Days:        3,
		SubregionID: "58581a836630e24c44878fd4", // Santa Barbara, CA
	}

	cs, err := c.GetConditions(cq)
	if err != nil {
		fmt.Printf("Error fetching Conditions: %v\n", err)
		return
	}

	fmt.Printf("Conditions: %v\n", cs)
}

func getTides() {
	bu, err := url.Parse(surflinef.TidesBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	c := surflinef.Client{BaseURL: bu}

	tq := surflinef.TidesQuery{
		Days:   3,
		SpotID: "5842041f4e65fad6a7708814", // Rincon, CA
	}

	ts, err := c.GetTides(tq)
	if err != nil {
		fmt.Printf("Error fetching Tides: %v\n", err)
		return
	}

	fmt.Printf("Tides: %v\n", ts)
}

func getTaxonomy() {
	bu, err := url.Parse(surflinef.TaxonomyBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	c := surflinef.Client{BaseURL: bu}

	tq := surflinef.TaxonomyQuery{
		ID:       "58f7ed58dadb30820bb38f8b", // Ventura County, CA
		MaxDepth: 1,
		Type:     "taxonomy",
	}

	t, err := c.GetTaxonomy(tq)
	if err != nil {
		fmt.Printf("Error fetching Taxonomy: %v\n", err)
		return
	}

	fmt.Printf("Taxonomy: %v\n", t)
}

func getWave() {
	bu, err := url.Parse(surflinef.WaveBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	c := surflinef.Client{BaseURL: bu}

	wq := surflinef.WaveQuery{
		SpotID:        "5842041f4e65fad6a7708814",
		Days:          1,
		IntervalHours: 6,
		MaxHeights:    false,
	}

	t, err := c.GetWave(wq)
	if err != nil {
		fmt.Printf("Error fetching Wave: %v\n", err)
		return
	}

	fmt.Printf("Wave: %v\n", t)
}
