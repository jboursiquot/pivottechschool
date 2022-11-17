package main

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getAnalyzer() ReleaseAnalyzer {
	return NewAnalyzer()
}

func TestAnalyzer_GetReleaseStats(t *testing.T) {
	analyzer := getAnalyzer()
	actual, err := analyzer.GetReleaseStats()
	assert.Nil(t, err)

	expected := []ReleaseStats{
		{releaseID: "1b6453892473a467d07372d45eb05abc2031647a", minQueryTime: 86, avgQueryTime: 214.44, maxQueryTime: 339},
		{releaseID: "fe5dbbcea5ce7e2988b8c69bcfdfde8904aabc1f", minQueryTime: 89, avgQueryTime: 224.71, maxQueryTime: 340},
		{releaseID: "c1dfd96eea8cc2b62785275bca38ac261256e278", minQueryTime: 91, avgQueryTime: 211.33, maxQueryTime: 304},
		{releaseID: "da4b9237bacccdf19c0760cab7aec4a8359010b0", minQueryTime: 86, avgQueryTime: 221.75, maxQueryTime: 339},
		{releaseID: "ac3478d69a3c81fa62e60f5c3696165a4e5e6ac4", minQueryTime: 86, avgQueryTime: 212.47, maxQueryTime: 340},
		{releaseID: "0ade7c2cf97f75d009975f4d720d1fa6c19f4897", minQueryTime: 100, avgQueryTime: 237.35, maxQueryTime: 336},
		{releaseID: "356a192b7913b04c54574d18c28d46e6395428ab", minQueryTime: 86, avgQueryTime: 215.32, maxQueryTime: 341},
		{releaseID: "77de68daecd823babbb58edb1c8e14d7106e83bb", minQueryTime: 88, avgQueryTime: 201.83, maxQueryTime: 336},
		{releaseID: "902ba3cda1883801594b6e1b452790cc53948fda", minQueryTime: 88, avgQueryTime: 208.94, maxQueryTime: 334}}

	sort.Slice(expected, func(i, j int) bool {
		return expected[i].releaseID > expected[j].releaseID
	})
	sort.Slice(actual, func(i, j int) bool {
		return actual[i].releaseID > actual[j].releaseID
	})

	assert.True(t, reflect.DeepEqual(expected, actual))
}

func TestAnalyzer_GetReleaseQuality(t *testing.T) {
	analyzer := getAnalyzer()
	actual, err := analyzer.GetReleaseQuality()
	assert.Nil(t, err)

	expected := ReleaseQuality{
		bestReleaseID:  "77de68daecd823babbb58edb1c8e14d7106e83bb",
		worstReleaseID: "0ade7c2cf97f75d009975f4d720d1fa6c19f4897",
	}
	assert.True(t, reflect.DeepEqual(expected, actual))
}

func TestAnalyzer_GetReleaseHistory(t *testing.T) {
	analyzer := getAnalyzer()
	actual, err := analyzer.GetReleaseHistory()
	assert.Nil(t, err)

	expected := []string{
		"356a192b7913b04c54574d18c28d46e6395428ab",
		"da4b9237bacccdf19c0760cab7aec4a8359010b0",
		"77de68daecd823babbb58edb1c8e14d7106e83bb",
		"1b6453892473a467d07372d45eb05abc2031647a",
		"ac3478d69a3c81fa62e60f5c3696165a4e5e6ac4",
		"c1dfd96eea8cc2b62785275bca38ac261256e278",
		"902ba3cda1883801594b6e1b452790cc53948fda",
		"fe5dbbcea5ce7e2988b8c69bcfdfde8904aabc1f",
		"0ade7c2cf97f75d009975f4d720d1fa6c19f4897"}

	sort.Slice(expected, func(i, j int) bool {
		return expected[i] > expected[j]
	})
	sort.Slice(actual, func(i, j int) bool {
		return actual[i] > actual[j]
	})
	assert.True(t, reflect.DeepEqual(expected, actual))
}
