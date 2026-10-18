package game

import (
	"errors"
	"fmt"
	"strings"
)

type Card struct {
	id     int
	name   string
	rarity CardRarity
	stats  Stats
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
	MetricIntelligence Metric = iota
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

func (o Outcome) String() string {
	switch o {
	case OutcomeDraw:
		return "Draw"
	case OutcomeFirstWins:
		return "First Wins"
	case OutcomeSecondWins:
		return "Second Wins"
	default:
		return "Unknown"
	}
}

func (c Card) CompareBy(other Card, m Metric) Outcome {
	a, b := c.Stats(m), other.Stats(m)
	switch {
	case a > b:
		return OutcomeFirstWins
	case a < b:
		return OutcomeSecondWins
	default:
		return OutcomeDraw
	}
}

func NewCard(name string, rarity CardRarity, s Stats) (Card, error) {
	var errs []error

	if strings.TrimSpace(name) == "" {
		errs = append(errs, errors.New("empty name"))
	}

	if len(s) < 4 {
		errs = append(errs, errors.New("too few stats"))
	}

	// TODO: допилить
	//if errs != nil {
	//	return (nil, errs)
	//}

}

func (c Card) Stats(metric Metric) int {
	return c.stats[metric]
}

func (c *Card) SetMetric(metric Metric, value int) bool {
	if value <= 0 {
		return false
	}

	c.stats[metric] = value
	return true
}

func (c Card) Rarity() CardRarity {
	return c.rarity
}

func (c *Card) SetRarity(rarity CardRarity) {
	c.rarity = rarity
}

func (c Card) Name() string {
	return c.name
}

func (c *Card) SetName(name string) {
	c.name = name
}
