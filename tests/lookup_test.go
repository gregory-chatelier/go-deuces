package deuces_test

import (
	"go-deuces"
	"testing"
)

func TestLookupTable_FlushLookupSize(t *testing.T) {
	lt := deuces.NewLookupTable()
	if len(lt.FlushLookup) != 1287 {
		t.Errorf("FlushLookup size = %d, want 1287", len(lt.FlushLookup))
	}
}

func TestLookupTable_UnsuitedLookupSize(t *testing.T) {
	lt := deuces.NewLookupTable()
	if len(lt.UnsuitedLookup) != 6175 {
		t.Errorf("UnsuitedLookup size = %d, want 6175", len(lt.UnsuitedLookup))
	}
}