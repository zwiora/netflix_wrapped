package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// --- parseTime ---

func TestParseTime_Valid(t *testing.T) {
	dur, err := parseTime("01:30:15")
	assert.NoError(t, err)
	assert.Equal(t, time.Hour+30*time.Minute+15*time.Second, dur)
}

func TestParseTime_Invalid(t *testing.T) {
	_, err := parseTime("99:99")
	assert.Error(t, err)
}

// --- parseMonth ---

func TestParseMonth_Valid1(t *testing.T) {
	month, err := parseMonth("1/2/2024 15:04")
	assert.NoError(t, err)
	assert.Equal(t, "2024.1", month)
}

func TestParseMonth_Valid2(t *testing.T) {
	month, err := parseMonth("2023-12-11 12:45:00")
	assert.NoError(t, err)
	assert.Equal(t, "2023.12", month)
}

func TestParseMonth_Invalid(t *testing.T) {
	_, err := parseMonth("not a date")
	assert.Error(t, err)
}

// --- IsWithinThreeHours ---

func TestIsWithinThreeHours_True(t *testing.T) {
	res, err := IsWithinThreeHours("1/2/2023 12:00", "1/2/2023 13:59")
	assert.NoError(t, err)
	assert.True(t, res)
}

func TestIsWithinThreeHours_False(t *testing.T) {
	res, err := IsWithinThreeHours("1/2/2023 12:00", "1/2/2023 16:01")
	assert.NoError(t, err)
	assert.False(t, res)
}

func TestIsWithinThreeHours_Invalid(t *testing.T) {
	_, err := IsWithinThreeHours("invalid", "1/2/2023 12:00")
	assert.Error(t, err)
}
