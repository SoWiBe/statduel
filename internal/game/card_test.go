package game

import "testing"

func TestCard_CompareBy(t *testing.T) {
	tests := []struct {
		name  string
		m     Metric
		c     Card
		other Card
		want  Outcome
	}{
		{"first is winner",
			MetricIntelligence,
			Card{Stats: [5]int{499}},
			Card{Stats: [5]int{250}},
			OutcomeFirstWins,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.CompareBy(tt.other, tt.m)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
