package gameassets

import (
	"testing"
	"github.com/gonza56d/go_snake/internal/gameassets"
)

func TestGenerateNewMapSuccess(t *testing.T) {
	tests := []struct {
		mapSizeOption uint8
		expected      [2]uint8
	}{
		{0, [2]uint8{25, 35}},
		{1, [2]uint8{40, 50}},
		{2, [2]uint8{55, 65}},
		{3, [2]uint8{75, 85}},
	}

	for _, test := range tests {
		result := gameassets.GenerateNewMap(test.mapSizeOption)

		if result.XSize < test.expected[0] || result.XSize > test.expected[1] ||
			result.YSize < test.expected[0] || result.YSize > test.expected[1] {
			t.Errorf("For mapSizeOption %d, expected XSize and YSize to have values between %d and %d. \n" +
				"(Got XSize=%d and YSize=%d)", test.mapSizeOption, test.expected[0], test.expected[1],
				result.XSize, result.YSize)
		}
	}
}

