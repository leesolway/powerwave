package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type UtilsTestSuite struct {
	suite.Suite
}

func (suite *UtilsTestSuite) TestDaysInMonth() {
	testCases := []struct {
		name     string
		input    time.Time
		expected int
	}{
		{"January", time.Date(2024, time.January, 15, 0, 0, 0, 0, time.UTC), 31},
		{"February (non-leap year)", time.Date(2023, time.February, 15, 0, 0, 0, 0, time.UTC), 28},
		{"February (leap year)", time.Date(2024, time.February, 15, 0, 0, 0, 0, time.UTC), 29},
		{"March", time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC), 31},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			actual := DaysInMonth(tc.input)
			suite.Equal(tc.expected, actual)
		})
	}
}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}
