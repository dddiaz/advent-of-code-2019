package main

import "testing"

func TestModuleFule(t *testing.T) {
	tables := []struct {
		weight        int
		expected_fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, table := range tables {
		result := ModuleFuel(table.weight)
		if result != table.expected_fuel {
			t.Errorf("ModuleFuel of (%d) was incorrect, got: %d, want: %d.", table.weight, result, table.expected_fuel)
		}
	}
}
