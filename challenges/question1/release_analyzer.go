package main

import (
	"encoding/json"
	"math"
	"os"
	"sort"
)

type ReleaseAnalyzer interface {
	GetReleaseStats() ([]ReleaseStats, error)
	GetReleaseQuality() (ReleaseQuality, error)
	GetReleaseHistory() ([]string, error)
}

type ReleaseStats struct {
	releaseID    string
	minQueryTime float64
	avgQueryTime float64
	maxQueryTime float64
}

type ReleaseQuality struct {
	bestReleaseID  string
	worstReleaseID string
}

type Event struct {
	Timestamp int     `json:"timestamp"`
	Version   string  `json:"version"`
	QueryTime float64 `json:"query_time"`
}

type Analyzer struct {
	releaseStats []ReleaseStats
	events       []Event
}

func NewAnalyzer() *Analyzer {
	a := Analyzer{}
	return &a
}

func (a Analyzer) roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func (a *Analyzer) analyze() error {
	if len(a.releaseStats) > 0 {
		return nil
	}

	f, err := os.Open("events.json")
	if err != nil {
		return err
	}
	if err := json.NewDecoder(f).Decode(&a.events); err != nil {
		return err
	}

	releaseStats := make(map[string]ReleaseStats)
	releaseEventCount := make(map[string]int)
	releaseTotalTime := make(map[string]float64)

	for _, event := range a.events {
		if _, ok := releaseStats[event.Version]; !ok {
			releaseStats[event.Version] = ReleaseStats{
				releaseID:    event.Version,
				minQueryTime: event.QueryTime,
				avgQueryTime: event.QueryTime,
				maxQueryTime: event.QueryTime,
			}
		} else {
			rs := releaseStats[event.Version]
			if event.QueryTime < rs.minQueryTime {
				rs.minQueryTime = event.QueryTime
			}
			if event.QueryTime > rs.maxQueryTime {
				rs.maxQueryTime = event.QueryTime
			}
			releaseStats[event.Version] = rs
		}
		releaseTotalTime[event.Version] += event.QueryTime
		releaseEventCount[event.Version]++
	}

	var stats []ReleaseStats
	for _, stat := range releaseStats {
		stat.avgQueryTime = a.roundFloat(releaseTotalTime[stat.releaseID]/float64(releaseEventCount[stat.releaseID]), 2)
		stats = append(stats, stat)
	}
	a.releaseStats = stats
	return nil
}

func (a *Analyzer) GetReleaseStats() ([]ReleaseStats, error) {
	if err := a.analyze(); err != nil {
		return nil, err
	}
	return a.releaseStats, nil
}

func (a *Analyzer) GetReleaseQuality() (ReleaseQuality, error) {
	if err := a.analyze(); err != nil {
		return ReleaseQuality{}, err
	}

	sort.Slice(a.releaseStats, func(i, j int) bool {
		return a.releaseStats[i].avgQueryTime < a.releaseStats[j].avgQueryTime
	})

	return ReleaseQuality{
		bestReleaseID:  a.releaseStats[0].releaseID,
		worstReleaseID: a.releaseStats[len(a.releaseStats)-1].releaseID,
	}, nil
}

func (a *Analyzer) GetReleaseHistory() ([]string, error) {
	if err := a.analyze(); err != nil {
		return nil, err
	}

	sort.Slice(a.events, func(i, j int) bool {
		return a.events[i].Timestamp < a.events[j].Timestamp
	})

	versionMap := make(map[string]struct{})
	for _, e := range a.events {
		versionMap[e.Version] = struct{}{}
	}

	keys := make([]string, 0, len(versionMap))
	for k := range versionMap {
		keys = append(keys, k)
	}

	return keys, nil
}
