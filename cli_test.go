package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDestinations(t *testing.T) {

	destinations := ParseDestinations(" LON , PAR   ,    , MAD")
	assert.Equal(t, 3, len(destinations))
	assert.Equal(t, "LON", destinations[0])
	assert.Equal(t, "PAR", destinations[1])
	assert.Equal(t, "MAD", destinations[2])

	expected := []string{"LON", "PAR", "MAD"}

	assert.Equal(t, expected, destinations)
}

func TestParseTimeOfDay(t *testing.T) {
	testParseTimeOfDay(t, 30, "0030")
	testParseTimeOfDay(t, 90, "0130")
	testParseTimeOfDay(t, 12*60+30, "1230")
	testParseTimeOfDay(t, 18*60, "1800")
}

func testParseTimeOfDay(t *testing.T, expected int, input string) {
	duration, err := ParseTimeOfDay(input)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, time.Minute*time.Duration(expected), duration)
}
