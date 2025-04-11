package clockface

import (
	"math"
	"testing"
	"time"
)

type radModel struct {
	time  time.Time
	angle float64
}
type pointCase struct {
	time  time.Time
	point Point
}

func TestSecondsInRadians(t *testing.T) {

	cases := []radModel{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}
	radiansTests(t, cases, SecondsInRadians)
}

func TestMinutesInRadians(t *testing.T) {

	cases := []radModel{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}
	radiansTests(t, cases, MinutesInRadians)
}
func TestHoursInRadians(t *testing.T) {

	cases := []radModel{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}
	radiansTests(t, cases, HoursInRadians)
}

func TestSecondHandPoint(t *testing.T) {
	cases := []pointCase{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	pointTests(t, cases, SecondHandPoint)
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []pointCase{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	pointTests(t, cases, MinuteHandPoint)
}

func TestHourHandPoint(t *testing.T) {
	cases := []pointCase{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	pointTests(t, cases, HourHandPoint)
}

func radiansTests(t *testing.T, cases []radModel, toRad func(time.Time) float64) {
	t.Helper()
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := toRad(c.time)

			if got != c.angle {
				t.Fatalf("Wanted %v radians but got %v", c.angle, got)
			}
		})
	}
}

func pointTests(t *testing.T, cases []pointCase, toPoint func(t time.Time) Point) {
	t.Helper()
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := toPoint(c.time)

			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityTreshold = 1e-7
	return math.Abs(a-b) < equalityTreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.Local)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
