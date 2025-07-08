package deuces_test

import (
	"go-deuces"
	"testing"
)

func TestNewCard(t *testing.T) {
	tests := []struct {
		input    string
		expected deuces.Card
		hasError bool
	}{
		{"As", deuces.Card(268442665), false},
		{"2c", deuces.Card(98306), false},
		{"Xx", deuces.Card(0), true},
	}

	for _, test := range tests {
		card, err := deuces.NewCard(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("NewCard(%q) error = %v, wantErr %v", test.input, err, test.hasError)
			continue
		}
		if card != test.expected {
			t.Errorf("NewCard(%q) = %v, want %v", test.input, card, test.expected)
		}
	}
}

func TestCard_IntToPrettyStr(t *testing.T) {
	card, _ := deuces.NewCard("As")
	expected := "A♠"
	if got := card.IntToPrettyStr(); got != expected {
		t.Errorf("IntToPrettyStr() = %q, want %q", got, expected)
	}
}

func TestCard_Getters(t *testing.T) {
	card, _ := deuces.NewCard("As")
	if rank := card.GetRankInt(); rank != 12 {
		t.Errorf("GetRankInt() = %d, want 12", rank)
	}
	if suit := card.GetSuitInt(); suit != 1 {
		t.Errorf("GetSuitInt() = %d, want 1", suit)
	}
	if bitrank := card.GetBitrankInt(); bitrank != 4096 {
		t.Errorf("GetBitrankInt() = %d, want 4096", bitrank)
	}
	if prime := card.GetPrime(); prime != 41 {
		t.Errorf("GetPrime() = %d, want 41", prime)
	}
}
