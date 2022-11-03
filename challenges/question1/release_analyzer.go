package main

import "errors"

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

type Analyzer struct{}

func NewAnalyzer() *Analyzer {
	a := Analyzer{}
	return &a
}

func (a *Analyzer) GetReleaseStats() ([]ReleaseStats, error) {
	return []ReleaseStats{}, errors.New("not implemented")
}

func (a *Analyzer) GetReleaseQuality() (ReleaseQuality, error) {
	return ReleaseQuality{}, errors.New("not implemented")
}

func (a *Analyzer) GetReleaseHistory() ([]string, error) {
	return nil, errors.New("not implemented")
}
