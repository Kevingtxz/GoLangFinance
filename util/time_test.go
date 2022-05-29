package util

import "testing"

func TestStringToDate(t *testing.T) {
	var convertedTime = StringToDate("2022-05-28T10:00:00")

	if convertedTime.Year() != 2022 {
		t.Errorf("Wait that the year is equal 2022")
	}

	if convertedTime.Month() != 05 {
		t.Errorf("Wait that the month is equal 05")
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("Wait that the hour is equal 10")
	}
}
