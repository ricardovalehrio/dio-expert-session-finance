package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2026-04-26T18:00:05")

	if convertedTime.Year() != 2026 {
		t.Errorf("Espero que o ano seja 2026")
	}
	if convertedTime.Month() != 04 {
		t.Errorf("Espero que o mês seja 04")
	}
	if convertedTime.Day() != 26 {
		t.Errorf("Espero que o dia seja 26")
	}
	if convertedTime.Hour() != 18 {
		t.Errorf("Espero que a hora seja 18")
	}
}
