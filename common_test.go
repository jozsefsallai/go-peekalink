package peekalink_test

import (
	"testing"

	"github.com/jozsefsallai/go-peekalink"
)

func TestDateString(t *testing.T) {
	t.Run("DateString#ToDate", func(t *testing.T) {
		t.Run("should parse date correctly", func(t *testing.T) {
			// Thu Feb 25 2021 20:01:37 GMT
			input := "2021-02-25T20:01:37.213993Z"

			result, err := peekalink.DateString(input).ToDate()
			if err != nil {
				t.Fatal(err)
			}

			year, month, day := result.Date()
			hour, minute, second := result.Hour(), result.Minute(), result.Second()

			if year != 2021 {
				t.Errorf("DateString#ToDate `year` got %v, want %v", year, 2021)
			}

			if month != 2 {
				t.Errorf("DateString#ToDate `month` got %v, want %v", month, 2)
			}

			if day != 25 {
				t.Errorf("DateString#ToDate `day` got %v, want %v", day, 25)
			}

			if hour != 20 {
				t.Errorf("DateString#ToDate `hour` got %v, want %v", hour, 20)
			}

			if minute != 1 {
				t.Errorf("DateString#ToDate `minute` got %v, want %v", minute, 1)
			}

			if second != 37 {
				t.Errorf("DateString#ToDate `second` got %v, want %v", second, 37)
			}
		})
	})
}
