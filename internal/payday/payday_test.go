package payday

import (
	"testing"
	"time"
)

func Test_daysIn(t *testing.T) {
	type args struct {
		m    time.Month
		year int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"april 2019", args{4, 2019}, 30},
		{"february 2017", args{2, 2017}, 28},
		{"february 2024", args{2, 2024}, 29},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := daysIn(tt.args.m, tt.args.year); got != tt.want {
				t.Errorf("daysIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LastFridayOf(t *testing.T) {

	t.Run("december 2018", func(t *testing.T) {
		result := LastFridayOf(12, 2018)
		expected := 28

		if result != expected {
			t.Errorf("got %d, expected %d", result, expected)
		}
	})

	t.Run("august 2019", func(t *testing.T) {
		result := LastFridayOf(8, 2019)
		expected := 30

		if result != expected {
			t.Errorf("got %d, expected %d", result, expected)
		}
	})

	t.Run("february 2019", func(t *testing.T) {
		result := LastFridayOf(2, 2019)
		expected := 22

		if result != expected {
			t.Errorf("got %d, expected %d", result, expected)
		}
	})
}

func TestWeeksLong(t *testing.T) {
	type args struct {
		m    time.Month
		year int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"april 2019", args{4, 2019}, 4},
		{"november 2018", args{11, 2018}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeeksLong(tt.args.m, tt.args.year); got != tt.want {
				t.Errorf("WeeksLong() = %v, want %v", got, tt.want)
			}
		})
	}
}
