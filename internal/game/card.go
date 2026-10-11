package game

import (
	"errors"
	"fmt"
)

type Card struct {
	ID     int
	Name   string
	Rarity CardRarity
	Stats  Stats
}

type CardRarity int

const (
	RarityBasic CardRarity = iota
	RarityRare
	RarityUltraRare
	RarityBlessing
)

type Metric int

const (
	MetricIntensity Metric = iota
	MetricDexterity
	MetricLuck
	MetricStrength
	MetricCombat
	metricCount // снаружи пакета не видна
)

type Outcome int

const (
	OutcomeDraw Outcome = iota
	OutcomeFirstWins
	OutcomeSecondWins
)

type Stats [metricCount]int

var ErrUnknownMetric = errors.New("unknown metric")

func ParseMetric(n int) (Metric, error) {
	if n < 0 || n >= int(metricCount) {
		return 0, fmt.Errorf("%w: %d", ErrUnknownMetric, n)
	}
	return Metric(n), nil
}

func (c Card) CompareBy(other Card, m Metric) Outcome {
	a, b := c.Stats[m], other.Stats[m]
	switch {
	case a < b:
		return OutcomeFirstWins
	case a > b:
		return OutcomeSecondWins
	default:
		return OutcomeDraw
	}
}
