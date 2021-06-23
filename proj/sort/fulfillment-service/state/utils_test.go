package state

import (
	"testing"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
	"github.com/stretchr/testify/assert"
)

func TestItemsEqual(t *testing.T) {
	testcases := []struct {
		name          string
		item1         *gen.Item
		item2         *gen.Item
		output        bool
		expectedError bool
	}{
		{
			name:          "test itemsEqual: no difference",
			item1:         &gen.Item{Code: "123", Label: "Coke"},
			item2:         &gen.Item{Code: "123", Label: "Coke"},
			output:        true,
			expectedError: false,
		},
		{
			name:          "test itemsEqual: code difference",
			item1:         &gen.Item{Code: "123", Label: "Coke"},
			item2:         &gen.Item{Code: "124", Label: "Coke"},
			output:        false,
			expectedError: false,
		},
		{
			name:          "test itemsEqual: label difference",
			item1:         &gen.Item{Code: "123", Label: "Coke"},
			item2:         &gen.Item{Code: "123", Label: "Other coke"},
			output:        false,
			expectedError: false,
		},
		{
			name:          "test itemsEqual: code and label difference",
			item1:         &gen.Item{Code: "123", Label: "Coke"},
			item2:         &gen.Item{Code: "124", Label: "Other coke"},
			output:        false,
			expectedError: false,
		},
		{
			name:          "test itemsEqual error: item1 is nil",
			item1:         nil,
			item2:         &gen.Item{Code: "124", Label: "Coke"},
			output:        false,
			expectedError: true,
		},
		{
			name:          "test itemsEqual error: item2  is nil",
			item1:         &gen.Item{Code: "123", Label: "Coke"},
			item2:         nil,
			output:        false,
			expectedError: true,
		},
		{
			name:          "test itemsEqual error: both items nil",
			item1:         nil,
			item2:         nil,
			output:        false,
			expectedError: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			itemsEqual, err := itemsEqual(tc.item1, tc.item2)

			if tc.output {
				assert.True(t, itemsEqual)
			} else {
				assert.False(t, itemsEqual)
			}
			if tc.expectedError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestIsEmptyItem(t *testing.T) {
	testcases := []struct {
		name          string
		item          *gen.Item
		output        bool
		expectedError bool
	}{
		{
			name:          "item with no code",
			item:          &gen.Item{Label: "coke"},
			output:        true,
			expectedError: false,
		},
		{
			name:          "item with no label",
			item:          &gen.Item{Code: "123"},
			output:        true,
			expectedError: false,
		},
		{
			name:          "item with no code and no label",
			item:          &gen.Item{},
			output:        true,
			expectedError: false,
		},
		{
			name:          "non-empty item",
			item:          &gen.Item{Code: "123", Label: "Coke"},
			output:        false,
			expectedError: false,
		},
		{
			name:          "nil item",
			item:          nil,
			output:        true,
			expectedError: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			isEmptyItem, err := isEmptyItem(tc.item)
			if tc.output {
				assert.True(t, isEmptyItem)
			} else {
				assert.False(t, isEmptyItem)
			}
			if tc.expectedError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
